package repository

import (
	"context"
	"github.com/Rafiur/wallet_app/internal/infrastructure/repository/schema"
)

type CurrencyRepository interface {
	Create(ctx context.Context, currency *schema.Currency) error
	GetByCode(ctx context.Context, code string) (*schema.Currency, error)
	List(ctx context.Context) ([]*schema.Currency, error)
	Update(ctx context.Context, currency *schema.Currency) error
	Delete(ctx context.Context, code string) error
}
