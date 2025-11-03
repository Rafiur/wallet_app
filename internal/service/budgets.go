package service

import (
	"context"
	"errors"
	"github.com/Rafiur/wallet_app/internal/infrastructure/repository"
	"github.com/Rafiur/wallet_app/internal/infrastructure/repository/schema"
)

type BudgetService struct {
	repo repository.BudgetRepository
}

func NewBudgetService(repo repository.BudgetRepository) *BudgetService {
	return &BudgetService{repo: repo}
}

func (svc *BudgetService) Create(ctx context.Context, budget *schema.Budget) error {
	if budget.UserID == "" || budget.Period == "" || budget.Amount <= 0 {
		return errors.New("user_id, period, and positive amount are required")
	}
	return svc.repo.Create(ctx, budget)
}

func (svc *BudgetService) GetByID(ctx context.Context, id string) (*schema.Budget, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return svc.repo.GetByID(ctx, id)
}

func (svc *BudgetService) GetByUserID(ctx context.Context, userID string) ([]*schema.Budget, error) {
	if userID == "" {
		return nil, errors.New("user_id is required")
	}
	return svc.repo.GetByUserID(ctx, userID)
}

func (svc *BudgetService) GetByCategoryAndPeriod(ctx context.Context, userID, categoryID, period string) (*schema.Budget, error) {
	if userID == "" || period == "" {
		return nil, errors.New("user_id and period are required")
	}
	return svc.repo.GetByCategoryAndPeriod(ctx, userID, categoryID, period)
}

func (svc *BudgetService) Update(ctx context.Context, budget *schema.Budget) error {
	if budget.ID == "" {
		return errors.New("id is required")
	}
	return svc.repo.Update(ctx, budget)
}

func (svc *BudgetService) Delete(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("id is required")
	}
	return svc.repo.Delete(ctx, id)
}
