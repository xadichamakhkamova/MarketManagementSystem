package service

import (
	"context"
	"debt-service/genproto/debtpb"
	"debt-service/internal/repository"
)

type DebtService struct {
	debtpb.UnimplementedDebtServiceServer
	repo repository.IDebtRepository
}

func NewDebtService(repo repository.IDebtRepository) *DebtService {
	return &DebtService{
		repo: repo,
	}
}

func (s *DebtService) CreateDebt(ctx context.Context, req *debtpb.CreateDebtReq) (*debtpb.DebtResp, error) {
	return s.repo.CreateDebt(ctx, req)
}
func (s *DebtService) UpdateDebt(ctx context.Context, req *debtpb.UpdateDebtReq) (*debtpb.DebtResp, error) {
	return s.repo.UpdateDebt(ctx, req)
}
func (s *DebtService) DeleteDebt(ctx context.Context, req *debtpb.DeleteDebtReq) (*debtpb.DebtResp, error) {
	return s.repo.DeleteDebt(ctx, req)
}
func (s *DebtService) GetDebtById(ctx context.Context, req *debtpb.GetDebtByIdReq) (*debtpb.DebtResp, error) {
	return s.repo.GetDebtById(ctx, req)
}
func (s *DebtService) GetDebtByFilter(ctx context.Context, req *debtpb.GetDebtByFilterReq) (*debtpb.GetDebtByFilterResp, error) {
	return s.repo.GetDebtByFilter(ctx, req)
}
