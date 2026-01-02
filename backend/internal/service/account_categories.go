package service

import (
	"context"
	"errors"
	"github.com/Rafiur/wallet_app/internal/infrastructure/repository"
	"github.com/Rafiur/wallet_app/internal/infrastructure/repository/schema"
)

type AccountCurrenciesService struct {
	repo repository.AccountCurrenciesRepository
}

func NewAccountCurrenciesService(repo repository.AccountCurrenciesRepository) *AccountCurrenciesService {
	return &AccountCurrenciesService{repo: repo}
}

func (svc *AccountCurrenciesService) Create(ctx context.Context, ac *schema.AccountCurrencies) error {
	if ac.CurrencyCode == "" || ac.AccountID == "" {
		return errors.New("currency_code and account_id are required")
	}
	return svc.repo.Create(ctx, ac)
}

func (svc *AccountCurrenciesService) GetByID(ctx context.Context, id string) (*schema.AccountCurrencies, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return svc.repo.GetByID(ctx, id)
}

func (svc *AccountCurrenciesService) GetByAccountID(ctx context.Context, accountID string) ([]*schema.AccountCurrencies, error) {
	if accountID == "" {
		return nil, errors.New("account_id is required")
	}
	return svc.repo.GetByAccountID(ctx, accountID)
}

func (svc *AccountCurrenciesService) Update(ctx context.Context, ac *schema.AccountCurrencies) error {
	if ac.ID == "" {
		return errors.New("id is required")
	}
	return svc.repo.Update(ctx, ac)
}

func (svc *AccountCurrenciesService) Delete(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("id is required")
	}
	return svc.repo.Delete(ctx, id)
}
