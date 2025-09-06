package service

import (
	"context"
	pb "dashboard-service/genproto/dashboardpb"
	"dashboard-service/internal/repository"
)

type DashboardService struct {
	pb.UnimplementedDashboardServiceServer
	reportRepo repository.IDashboardRepository
}

func NewDashboardService(repo repository.IDashboardRepository) *DashboardService {
	return &DashboardService{
		reportRepo: repo,
	}
}

func (service *DashboardService) GetDashboardReport(ctx context.Context, req *pb.GetDashboardReportRequest) (*pb.GetDashboardReportResponse, error) {
	return service.reportRepo.GetDashboardReport()
}