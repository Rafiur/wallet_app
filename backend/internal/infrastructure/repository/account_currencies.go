package repository

import (
	"context"
	"github.com/Rafiur/wallet_app/internal/infrastructure/repository/schema"
)

type AccountCurrenciesRepository interface {
	Create(ctx context.Context, ac *schema.AccountCurrencies) error
	GetByID(ctx context.Context, id string) (*schema.AccountCurrencies, error)
	GetByAccountID(ctx context.Context, accountID string) ([]*schema.AccountCurrencies, error)
	Update(ctx context.Context, ac *schema.AccountCurrencies) error
	Delete(ctx context.Context, id string) error
}
