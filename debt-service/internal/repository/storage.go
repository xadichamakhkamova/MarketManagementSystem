package repository

import (
	"context"
	"database/sql"
	"debt-service/genproto/debtpb"
	"debt-service/internal/storage"
	"debt-service/logger"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type DebtREPO struct {
	queries *storage.Queries
}

func NewDebtSqlc(db *sql.DB) *storage.Queries {
	return storage.New(db)
}

func (q *DebtREPO) CreateDebt(ctx context.Context, req *debtpb.CreateDebtReq) (*debtpb.DebtResp, error) {
	logger.Info("CreateDebt: Started for ", req.FirstName, " ", req.LastName)
	
	// 02.01.2006 -> yy.mm.dd time format
	deadlineTime, err := time.Parse("2006-01-02", req.Deadline)
	if err != nil {
		logger.Error("CreateDebt: Invalid deadline format - ", err)
		return nil, err
	}
	
	
	floatPricePaid, err := strconv.ParseFloat(req.PricePaid, 64)
	if err != nil {
		logger.Error("CreateDebt: Invalid price paid - ", err)
		return nil, err
	}

	priceFloat, err := strconv.ParseFloat(req.Price, 64)
	if err != nil {
		logger.Error("CreateDebt: Invalid price - ", err)
		return nil, err
	}

	resp, err := q.queries.CreateDebt(ctx, storage.CreateDebtParams{
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		PhoneNumber:  req.PhoneNumber,
		Jshshir:      req.Jshshir,
		Address:      req.Address,
		BagID:        req.BagId,
		Price:        priceFloat,
		PricePaid:    floatPricePaid,
		Acquaintance: sql.NullString{String: req.Acquaintance, Valid: req.Acquaintance != ""},
		Collateral:   sql.NullString{String: req.Collateral, Valid: req.Collateral != ""},
		Deadline:     deadlineTime,
	})

	if err != nil {
		logger.Error("CreateDebt: Database error - ", err)
		return nil, err
	}

	logger.Info("CreateDebt: Successfully created debt for ", req.FirstName, " ", req.LastName)

	return &debtpb.DebtResp{
		Status:  true,
		Message: "Debt Created Successfully",
		Debt: &debtpb.Debt{
			Id:           resp.ID.String(),
			FirstName:    resp.FirstName,
			LastName:     resp.LastName,
			PhoneNumber:  resp.PhoneNumber,
			Jshshir:      resp.Jshshir,
			Address:      resp.Address,
			Price:        req.Price,
			PricePaid:    req.PricePaid,
			Acquaintance: resp.Acquaintance.String,
			Collateral:   resp.Collateral.String,
			Deadline:     resp.Deadline.String(),
			DebtCUD: &debtpb.DebtCUD{
				CreatedAt: resp.CreatedAt.Time.String(),
				UpdatedAt: resp.UpdatedAt.Time.String(),
				DeletedAt: int64(resp.DeletedAt.Int64),
			},
		},
	}, nil
}

func (q *DebtREPO) UpdateDebt(ctx context.Context, req *debtpb.UpdateDebtReq) (*debtpb.DebtResp, error) {
	logger.Info("UpdateDebt: Started for debt ID ", req.Id)
	
	// 2006.01.02 -> yy.mm.dd time format
	deadlineTime, err := time.Parse("2006-01-02", req.Deadline)
	if err != nil {
		logger.Error("UpdateDebt: Invalid deadline format - ", err)
		return nil, err
	}

	id, err := uuid.Parse(req.Id)
	if err != nil {
		logger.Error("UpdateDebt: Invalid UUID - ", err)
		return nil, err
	}
	
	floatPricePaid, err := strconv.ParseFloat(req.PricePaid, 64)
	if err != nil {
		logger.Error("UpdateDebt: Invalid price paid - ", err)
		return nil, err
	}
	
	resp, err := q.queries.UpdateDebt(ctx, storage.UpdateDebtParams{
		ID:           id,
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		PhoneNumber:  req.PhoneNumber,
		Jshshir:      req.Jshshir,
		Address:      req.Address,
		PricePaid:    floatPricePaid,
		Acquaintance: sql.NullString{String: req.Acquaintance, Valid: req.Acquaintance != ""},
		Collateral:   sql.NullString{String: req.Collateral, Valid: req.Collateral != ""},
		Deadline:     deadlineTime,
		UpdatedAt:    sql.NullTime{Time: time.Now()},
	})

	if err != nil {
		logger.Error("UpdateDebt: Database error - ", err)
		return nil, err
	}

	logger.Info("UpdateDebt: Successfully updated debt for ID ", req.Id)

	return &debtpb.DebtResp{
		Status:  true,
		Message: "Debt Updated Successfully",
		Debt: &debtpb.Debt{
			Id:           resp.ID.String(),
			FirstName:    resp.FirstName,
			LastName:     resp.LastName,
			PhoneNumber:  resp.PhoneNumber,
			Jshshir:      resp.Jshshir,
			Address:      resp.Address,
			PricePaid:    req.PricePaid,
			Acquaintance: resp.Acquaintance.String,
			Collateral:   resp.Collateral.String,
			Deadline:     resp.Deadline.String(),
			DebtCUD: &debtpb.DebtCUD{
				CreatedAt: resp.CreatedAt.Time.String(),
				UpdatedAt: resp.UpdatedAt.Time.String(),
				DeletedAt: int64(resp.DeletedAt.Int64),
			},
		},
	}, nil
}

func (q *DebtREPO) DeleteDebt(ctx context.Context, req *debtpb.DeleteDebtReq) (*debtpb.DebtResp, error) {
	logger.Info("DeleteDebt: Started for debt ID ", req.Id)

	id, err := uuid.Parse(req.Id)
	if err != nil {
		logger.Error("DeleteDebt: Invalid UUID - ", err)
		return nil, err
	}

	err = q.queries.DeleteDebt(ctx, storage.DeleteDebtParams{
		ID:        id,
		DeletedAt: sql.NullInt64{Int64: int64(time.Now().Unix())},
	})
	if err != nil {
		logger.Error("DeleteDebt: Database error - ", err)
		return nil, err
	}

	logger.Info("DeleteDebt: Successfully deleted debt with ID ", req.Id)

	return &debtpb.DebtResp{
		Status:  true,
		Message: "Debt Deleted Successfully",
	}, nil
}

func (q *DebtREPO) GetDebtById(ctx context.Context, req *debtpb.GetDebtByIdReq) (*debtpb.DebtResp, error) {
	logger.Info("GetDebtById: Started for debt ID ", req.Id)

	id, err := uuid.Parse(req.Id)
	if err != nil {
		logger.Error("GetDebtById: Invalid UUID - ", err)
		return nil, err
	}

	resp, err := q.queries.GetDebtById(ctx, id)
	if err != nil {
		logger.Error("GetDebtById: Database error - ", err)
		return nil, err
	}

	logger.Info("GetDebtById: Successfully retrieved debt for ID ", req.Id)

	return &debtpb.DebtResp{
		Status:  true,
		Message: "Get Debt Successfully",
		Debt: &debtpb.Debt{
			Id:           resp.ID.String(),
			FirstName:    resp.FirstName,
			LastName:     resp.LastName,
			PhoneNumber:  resp.PhoneNumber,
			Jshshir:      resp.Jshshir,
			Address:      resp.Address,
			PricePaid:    strconv.FormatFloat(resp.PricePaid, 'f', -1, 64),
			Acquaintance: resp.Acquaintance.String,
			Collateral:   resp.Collateral.String,
			Deadline:     resp.Deadline.String(),
			DebtCUD: &debtpb.DebtCUD{
				CreatedAt: resp.CreatedAt.Time.String(),
				UpdatedAt: resp.UpdatedAt.Time.String(),
				DeletedAt: int64(resp.DeletedAt.Int64),
			},
		},
	}, nil
}

func (q *DebtREPO) GetDebtByFilter(ctx context.Context, req *debtpb.GetDebtByFilterReq) (*debtpb.GetDebtByFilterResp, error) {
	logger.Info("GetDebtByFilter: Started with search parameter ", req.Search)

	resp, err := q.queries.GetDebtByFilter(ctx, req.Search)
	if err != nil {
		logger.Error("GetDebtByFilter: Database error - ", err)
		return nil, err
	}

	debts := make([]*debtpb.Debt, 0, len(resp))

	for _, i := range resp {
		debts = append(debts, &debtpb.Debt{
			Id:           i.ID.String(),
			FirstName:    i.FirstName,
			LastName:     i.LastName,
			PhoneNumber:  i.PhoneNumber,
			Jshshir:      i.Jshshir,
			Address:      i.Address,
			PricePaid:    strconv.FormatFloat(i.PricePaid, 'f', -1, 64),
			Acquaintance: i.Acquaintance.String,
			Collateral:   i.Collateral.String,
			Deadline:     i.Deadline.String(),
			DebtCUD: &debtpb.DebtCUD{
				CreatedAt: i.CreatedAt.Time.String(),
				UpdatedAt: i.UpdatedAt.Time.String(),
				DeletedAt: int64(i.DeletedAt.Int64),
			},
		})
	}

	logger.Info("GetDebtByFilter: Successfully retrieved ", len(resp), " debts")

	return &debtpb.GetDebtByFilterResp{
		Status:       true,
		Message:      "Get By Filter Successfully",
		GetCountResp: int32(len(resp)),
		Debt:         debts,
	}, nil
}