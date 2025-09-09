package repository

import (
	"context"
	"debt-service/genproto/debtpb"
	"debt-service/internal/storage"
)

func NewIDebtRepository(queries *storage.Queries) IDebtRepository {
	return &DebtREPO{
		queries: queries,
	}
}

type IDebtRepository interface {
	CreateDebt(ctx context.Context, req *debtpb.CreateDebtReq) (*debtpb.DebtResp, error)
	UpdateDebt(ctx context.Context, req *debtpb.UpdateDebtReq) (*debtpb.DebtResp, error)
	DeleteDebt(ctx context.Context, req *debtpb.DeleteDebtReq) (*debtpb.DebtResp, error)
	GetDebtById(ctx context.Context, req *debtpb.GetDebtByIdReq) (*debtpb.DebtResp, error)
	GetDebtByFilter(ctx context.Context, req *debtpb.GetDebtByFilterReq) (*debtpb.GetDebtByFilterResp, error)
}
