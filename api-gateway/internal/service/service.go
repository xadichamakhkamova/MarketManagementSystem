package service

import (
	pbDashboard "api-gateway/genproto/dashboardpb"
	pbDebt "api-gateway/genproto/debtpb"
	pbProduct "api-gateway/genproto/productpb"
	"context"
)

type ServiceRepositoryClient struct {
	productClient   pbProduct.ProductServiceClient
	dashboardClient pbDashboard.DashboardServiceClient
	debtClient      pbDebt.DebtServiceClient
}

func NewServiceRepositoryClient(
	conn1 *pbProduct.ProductServiceClient,
	conn2 *pbDashboard.DashboardServiceClient,
	conn3 *pbDebt.DebtServiceClient,
) *ServiceRepositoryClient {
	return &ServiceRepositoryClient{
		productClient:   *conn1,
		dashboardClient: *conn2,
		debtClient:      *conn3,
	}
}

// Product methods
func (s *ServiceRepositoryClient) CreateProduct(ctx context.Context, req *pbProduct.CreateProductRequest) (*pbProduct.CreateProductResponse, error) {
	return s.productClient.CreateProduct(ctx, req)
}

func (s *ServiceRepositoryClient) GetProductById(ctx context.Context, req *pbProduct.GetProductByIdRequest) (*pbProduct.GetProductByIdResponse, error) {
	return s.productClient.GetProductById(ctx, req)
}

func (s *ServiceRepositoryClient) GetProductByFilter(ctx context.Context, req *pbProduct.GetProductByFilterRequest) (*pbProduct.GetProductByFilterResponse, error) {
	return s.productClient.GetProductByFilter(ctx, req)
}

func (s *ServiceRepositoryClient) UpdateStock(ctx context.Context, req *pbProduct.UpdateStockRequest) (*pbProduct.UpdateStockResponse, error) {
	return s.productClient.UpdateStock(ctx, req)
}

func (s *ServiceRepositoryClient) UpdateProduct(ctx context.Context, req *pbProduct.UpdateProductRequest) (*pbProduct.UpdateProductResponse, error) {
	return s.productClient.UpdateProduct(ctx, req)
}

func (s *ServiceRepositoryClient) DeleteProduct(ctx context.Context, req *pbProduct.DeleteProductRequest) (*pbProduct.DeleteProductResponse, error) {
	return s.productClient.DeleteProduct(ctx, req)
}

// Dashboard methods
func (s *ServiceRepositoryClient) GetDashboardReport(ctx context.Context, req *pbDashboard.GetDashboardReportRequest) (*pbDashboard.GetDashboardReportResponse, error) {
	return s.dashboardClient.GetDashboardReport(ctx, req)
}

// Debt methods
func (s *ServiceRepositoryClient) CreateDebt(ctx context.Context, req *pbDebt.CreateDebtReq) (*pbDebt.DebtResp, error) {
	return s.debtClient.CreateDebt(ctx, req)
}
func (s *ServiceRepositoryClient) UpdateDebt(ctx context.Context, req *pbDebt.UpdateDebtReq) (*pbDebt.DebtResp, error) {
	return s.debtClient.UpdateDebt(ctx, req)
}
func (s *ServiceRepositoryClient) DeleteDebt(ctx context.Context, req *pbDebt.DeleteDebtReq) (*pbDebt.DebtResp, error) {
	return s.debtClient.DeleteDebt(ctx, req)
}
func (s *ServiceRepositoryClient) GetDebtById(ctx context.Context, req *pbDebt.GetDebtByIdReq) (*pbDebt.DebtResp, error) {
	return s.debtClient.GetDebtById(ctx, req)
}
func (s *ServiceRepositoryClient) GetDebtByFilter(ctx context.Context, req *pbDebt.GetDebtByFilterReq) (*pbDebt.GetDebtByFilterResp, error) {
	return s.debtClient.GetDebtByFilter(ctx, req)
}
