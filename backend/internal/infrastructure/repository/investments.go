package repository

import (
	"context"
	"github.com/Rafiur/wallet_app/internal/infrastructure/repository/schema"
)

type InvestmentRepository interface {
	Create(ctx context.Context, req *schema.Investment) error
	GetByID(ctx context.Context, id string) (*schema.Investment, error)
	GetByUserID(ctx context.Context, userID string) ([]*schema.Investment, error)
	Update(ctx context.Context, req *schema.Investment) error
	Delete(ctx context.Context, id string) error
}
