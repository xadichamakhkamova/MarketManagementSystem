package main

import (
	api "api-gateway/internal/https"
	config "api-gateway/internal/config"
	_ "api-gateway/docs"
	
	productService "api-gateway/internal/clients/product-service"
	dashboardService "api-gateway/internal/clients/dashboard-service"
	debtService "api-gateway/internal/clients/debt-service"
	service "api-gateway/internal/service"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"api-gateway/logger"
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	// Loggerni boshlash
	logger.InitLog()

	// Konfiguratsiya faylini yuklash
	cfg, err := config.Load("config/config.yaml")
	if err != nil {
		logger.Fatal("Failed to load configuration:", err)
	}
	logger.Info("Configuration loaded successfully")

	// Product Service bilan ulanish
	conn1, err := productService.DialWithProductService(*cfg)
	if err != nil {
		logger.Fatal("Failed to connect to Product Service:", err)
	}
	logger.Info("Connected to Product Service")

	// Dashboard Service bilan ulanish
	conn2, err := dashboardService.DialWithDashboardService(*cfg)
	if err != nil {
		logger.Fatal("Failed to connect to Product Service:", err)
	}
	logger.Info("Connected to Product Service")

	// Debt Service bilan ulanish
	conn3, err := debtService.DialWithDebtService(*cfg)
	if err != nil {
		logger.Fatal("Failed to connect to Debt Service:", err)
	}
	logger.Info("Connected to Debt Service")

	// Product Service uchun mijoz xizmatini yaratish
	clientService := service.NewServiceRepositoryClient(conn1, conn2, conn3)
	logger.Info("Service clients initialized")

	// MinIO bilan ulanish
	endpoint := fmt.Sprintf("%s:%d", cfg.Minio.Host, cfg.Minio.Port)
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.Minio.Id, cfg.Minio.Secret, ""),
		Secure: false,
	})
	if err != nil {
		logger.Fatal("MinIO connection error - ", err)
	}
	logger.Info("Connected to MinIO")

	
	// API Gateway serverini ishga tushirish
	srv := api.NewGin(clientService, cfg.ApiGateway.Port, minioClient, endpoint)
	addr := fmt.Sprintf(":%d", cfg.ApiGateway.Port)

	// Signalni kutish uchun kanal yaratish (SIGINT yoki SIGTERM)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		logger.Info("Starting API Gateway on: ", addr)
		if err := srv.ListenAndServeTLS(cfg.CertFile, cfg.KeyFile); err != nil {
			logger.Fatal(err)
		}
	}()
	logger.Info("Starting API Gateway on address:", addr)
	
	// Signalni qabul qilish
	signalReceived := <-sigChan
	logger.Info("Received signal:", signalReceived)
	
	// Xizmatni to'xtatish uchun kontekst yaratish
	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownRelease()

	// API Gatewayni to'xtatish
	if err := srv.Shutdown(shutdownCtx); err != nil {
		logger.Fatal("Server shutdown error: ", err)
	}
	logger.Info("Graceful shutdown complete.")
}
