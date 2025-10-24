package repository

import (
	"context"
	"github.com/Rafiur/wallet_app/internal/infrastructure/repository/schema"
)

type BankRepository interface {
	Create(ctx context.Context, bank *schema.Bank) error
	GetByID(ctx context.Context, id string) (*schema.Bank, error)
	GetByUserID(ctx context.Context, userID string) ([]*schema.Bank, error)
	GetByAccountID(ctx context.Context, accountID string) (*schema.Bank, error)
	Update(ctx context.Context, bank *schema.Bank) error
	Delete(ctx context.Context, id string) error
}
