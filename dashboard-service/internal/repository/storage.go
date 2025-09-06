package repository

import (
	"context"
	pb "dashboard-service/genproto/dashboardpb"
	"dashboard-service/internal/repository/models"
	"dashboard-service/internal/storage"
	"dashboard-service/logger"
	"database/sql"
	"fmt"
)

type DashboardRepo struct {
	storage *storage.Queries
}

func NewDashboardSqlc(db *sql.DB) *storage.Queries {
	return storage.New(db)
}

func (db *DashboardRepo) UpsertProductSales(req models.ProductSalesUpdateRequest) error {

	logger.Info("UpsertProductSales: started: ", req.ProductId)

	res, err := db.storage.UpsertProductSales(context.Background(), storage.UpsertProductSalesParams{
		ProductID:    req.ProductId,
		Color:        req.ProductColor,
		CostPrice:    req.CostPrice,
		SellingPrice: req.SellingPrice,
	})

	if err != nil {
		logger.Error("UpsertProductSales: error while upserting product sales ", req.ProductId, ":", err)
		return err
	}
	if res != "success" {
		logger.Warn("UpsertProductSales: no rows affected for product_id - ", req.ProductId)
		return fmt.Errorf("no rows affected")
	}

	logger.Info("UpsertProductSales: product sales upserted successfully for product_id - ", req.ProductId)
	return nil
}

func (db *DashboardRepo) GetDashboardReport() (*pb.GetDashboardReportResponse, error) {

	logger.Info("GetDashboardReport: started")

	res, err := db.storage.GetDashboardReport(context.Background())
	if err != nil {
		logger.Error("GetDashboardReport: error retrieving report - ", err)
		return nil, err
	}

	report := pb.GetDashboardReportResponse{
		TotalSales:     res.TotalSales,
		TotalCostPrice: res.TotalCostPrice,
		TotalNetProfit: res.TotalNetProfit,
		TotalUnitsSold: res.TotalUnitsSold,
	}

	logger.Info("GetDashboardReport: successfully retrieved report data")
	return &report, nil
}
