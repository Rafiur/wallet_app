package service

import (
	"context"
	"errors"
	"github.com/Rafiur/wallet_app/internal/infrastructure/repository"
	"github.com/Rafiur/wallet_app/internal/infrastructure/repository/schema"

	"time"
)

type CashFlowSummaryService struct {
	repo repository.CashFlowSummaryRepository
}

func NewCashFlowSummaryService(repo repository.CashFlowSummaryRepository) *CashFlowSummaryService {
	return &CashFlowSummaryService{repo: repo}
}

func (svc *CashFlowSummaryService) Create(ctx context.Context, cfs *schema.CashFlowSummary) error {
	if cfs.UserID == "" || cfs.Period == "" {
		return errors.New("user_id and period are required")
	}
	return svc.repo.Create(ctx, cfs)
}

func (svc *CashFlowSummaryService) GetByID(ctx context.Context, id string) (*schema.CashFlowSummary, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return svc.repo.GetByID(ctx, id)
}

func (svc *CashFlowSummaryService) GetByUserIDAndPeriod(ctx context.Context, userID, period string, startDate time.Time) ([]*schema.CashFlowSummary, error) {
	if userID == "" || period == "" {
		return nil, errors.New("user_id and period are required")
	}
	return svc.repo.GetByUserIDAndPeriod(ctx, userID, period, startDate)
}

func (svc *CashFlowSummaryService) Update(ctx context.Context, cfs *schema.CashFlowSummary) error {
	if cfs.ID == "" {
		return errors.New("id is required")
	}
	return svc.repo.Update(ctx, cfs)
}

func (svc *CashFlowSummaryService) Delete(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("id is required")
	}
	return svc.repo.Delete(ctx, id)
}
