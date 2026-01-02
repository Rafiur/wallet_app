package repository

import (
	"context"
	"github.com/Rafiur/wallet_app/internal/infrastructure/repository/schema"
)

type BudgetRepository interface {
	Create(ctx context.Context, budget *schema.Budget) error
	GetByID(ctx context.Context, id string) (*schema.Budget, error)
	GetByUserID(ctx context.Context, userID string) ([]*schema.Budget, error)
	GetByCategoryAndPeriod(ctx context.Context, userID, categoryID, period string) (*schema.Budget, error)
	Update(ctx context.Context, budget *schema.Budget) error
	Delete(ctx context.Context, id string) error
}
