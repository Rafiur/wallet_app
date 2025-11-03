package service

import (
	"context"
	"errors"
	"github.com/Rafiur/wallet_app/internal/infrastructure/repository"
	"github.com/Rafiur/wallet_app/internal/infrastructure/repository/schema"
)

type CurrencyService struct {
	repo repository.CurrencyRepository
}

func NewCurrencyService(repo repository.CurrencyRepository) *CurrencyService {
	return &CurrencyService{repo: repo}
}

func (svc *CurrencyService) Create(ctx context.Context, currency *schema.Currency) error {
	if currency.Code == "" || currency.Name == "" || currency.Symbol == "" {
		return errors.New("code, name, and symbol are required")
	}
	return svc.repo.Create(ctx, currency)
}

func (svc *CurrencyService) GetByCode(ctx context.Context, code string) (*schema.Currency, error) {
	if code == "" {
		return nil, errors.New("code is required")
	}
	return svc.repo.GetByCode(ctx, code)
}

func (svc *CurrencyService) List(ctx context.Context) ([]*schema.Currency, error) {
	return svc.repo.List(ctx)
}

func (svc *CurrencyService) Update(ctx context.Context, currency *schema.Currency) error {
	if currency.Code == "" {
		return errors.New("code is required")
	}
	return svc.repo.Update(ctx, currency)
}

func (svc *CurrencyService) Delete(ctx context.Context, code string) error {
	if code == "" {
		return errors.New("code is required")
	}
	return svc.repo.Delete(ctx, code)
}
