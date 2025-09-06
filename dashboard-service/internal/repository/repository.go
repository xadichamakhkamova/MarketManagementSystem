package repository

import (
	pb "dashboard-service/genproto/dashboardpb"
	"dashboard-service/internal/repository/models"
	"dashboard-service/internal/storage"
)

func NewIDashboardRepository(storage *storage.Queries) IDashboardRepository {
	return &DashboardRepo{
		storage: storage,
	}
}

type IDashboardRepository interface {
	UpsertProductSales(req models.ProductSalesUpdateRequest) error
	GetDashboardReport() (*pb.GetDashboardReportResponse, error)
}
