package repository

import (
	"context"
	"github.com/Rafiur/wallet_app/internal/infrastructure/repository/schema"

	"time"
)

type CashFlowSummaryRepository interface {
	Create(ctx context.Context, cfs *schema.CashFlowSummary) error
	GetByID(ctx context.Context, id string) (*schema.CashFlowSummary, error)
	GetByUserIDAndPeriod(ctx context.Context, userID, period string, startDate time.Time) ([]*schema.CashFlowSummary, error)
	Update(ctx context.Context, cfs *schema.CashFlowSummary) error
	Delete(ctx context.Context, id string) error
}
