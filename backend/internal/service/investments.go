package service

import (
	"context"
	"errors"
	"github.com/Rafiur/wallet_app/internal/infrastructure/repository"
	"github.com/Rafiur/wallet_app/internal/infrastructure/repository/schema"
)

type InvestmentService struct {
	repo repository.InvestmentRepository
}

func NewInvestmentService(repo repository.InvestmentRepository) *InvestmentService {
	return &InvestmentService{repo: repo}
}

func (svc *InvestmentService) Create(ctx context.Context, inv *schema.Investment) error {
	if inv.UserID == "" || inv.Type == "" || inv.Amount <= 0 {
		return errors.New("user_id, type, and positive amount are required")
	}
	return svc.repo.Create(ctx, inv)
}

func (svc *InvestmentService) GetByID(ctx context.Context, id string) (*schema.Investment, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return svc.repo.GetByID(ctx, id)
}

func (svc *InvestmentService) GetByUserID(ctx context.Context, userID string) ([]*schema.Investment, error) {
	if userID == "" {
		return nil, errors.New("user_id is required")
	}
	return svc.repo.GetByUserID(ctx, userID)
}

func (svc *InvestmentService) Update(ctx context.Context, inv *schema.Investment) error {
	if inv.ID == "" {
		return errors.New("id is required")
	}
	return svc.repo.Update(ctx, inv)
}

func (svc *InvestmentService) Delete(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("id is required")
	}
	return svc.repo.Delete(ctx, id)
}
