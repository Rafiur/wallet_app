package repository

import (
	"context"
	"github.com/Rafiur/wallet_app/internal/infrastructure/repository/schema"

	"time"
)

type RecurringTransactionRepository interface {
	Create(ctx context.Context, rt *schema.RecurringTransaction) error
	GetByID(ctx context.Context, id string) (*schema.RecurringTransaction, error)
	GetByUserID(ctx context.Context, userID string) ([]*schema.RecurringTransaction, error)
	GetDueByDate(ctx context.Context, dueDate time.Time) ([]*schema.RecurringTransaction, error)
	Update(ctx context.Context, rt *schema.RecurringTransaction) error
	Delete(ctx context.Context, id string) error
}
