package main

import (
	"context"
	app "dashboard-service/internal/app"
	config "dashboard-service/internal/config"
	kafka "dashboard-service/internal/kafka/consumer"
	"dashboard-service/internal/repository"
	pq "dashboard-service/internal/repository/postgres"
	service "dashboard-service/internal/service"
	"dashboard-service/logger"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {

	logger.InitLog()

	cfg, err := config.Load("./config/config.yaml")
	if err != nil {
		logger.Fatal(err)
	}
	logger.Info("Configuration loaded")

	db, err := pq.ConnectDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Info("Connected to Postgresql")

	queries := repository.NewDashboardSqlc(db)
	repo := repository.NewIDashboardRepository(queries)

	consumer, err := kafka.NewConsumeInit(cfg, repo)
	if err != nil {
		logger.Fatal(err)
	}
	defer consumer.Close()

	srv := service.NewDashboardService(repo)

	application := app.New(*srv, cfg.Service.Port)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		consumer.ConsumeMessage()
	}()

	go func() {
		defer wg.Done()
		application.MustRun()
	}()
	wg.Wait()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	sig := <-stop

	logger.Info("Received signal: ", sig)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	application.Stop()
	<-ctx.Done()
	logger.Info("Graceful shutdown complete.")
}
