package main

import (
	"context"
	app "debt-service/internal/app"
	config "debt-service/internal/config"
	"debt-service/internal/repository"
	pq "debt-service/internal/repository/postgres"
	service "debt-service/internal/service"
	"debt-service/logger"
	"os"
	"os/signal"
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

	queries := repository.NewDebtSqlc(db)
	repo := repository.NewIDebtRepository(queries)

	srv := service.NewDebtService(repo)

	application := app.New(*srv, cfg.Service.Port)

	go func() {
		application.MustRun()
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	sig := <-stop

	logger.Info("Received signal: ", sig)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	application.Stop()
	<-ctx.Done()
}
