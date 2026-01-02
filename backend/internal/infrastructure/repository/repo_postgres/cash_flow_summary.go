package repo_postgres

import (
	"context"
	"github.com/Rafiur/wallet_app/internal/infrastructure/repository/schema"

	"github.com/uptrace/bun"
	"time"
)

type CashFlowSummaryRepository struct {
	db *bun.DB
}

func NewCashFlowSummaryRepository(db *bun.DB) *CashFlowSummaryRepository {
	return &CashFlowSummaryRepository{db: db}
}

func (r *CashFlowSummaryRepository) Create(ctx context.Context, cfs *schema.CashFlowSummary) error {
	_, err := r.db.NewInsert().Model(cfs).Exec(ctx)
	return err
}

func (r *CashFlowSummaryRepository) GetByID(ctx context.Context, id string) (*schema.CashFlowSummary, error) {
	cfs := new(schema.CashFlowSummary)
	err := r.db.NewSelect().
		Model(cfs).
		Where("id = ?", id).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return cfs, nil
}

func (r *CashFlowSummaryRepository) GetByUserIDAndPeriod(ctx context.Context, userID, period string, startDate time.Time) ([]*schema.CashFlowSummary, error) {
	var cfs []*schema.CashFlowSummary
	err := r.db.NewSelect().
		Model(&cfs).
		Where("user_id = ?", userID).
		Where("period = ?", period).
		Where("start_date >= ?", startDate).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return cfs, nil
}

func (r *CashFlowSummaryRepository) Update(ctx context.Context, cfs *schema.CashFlowSummary) error {
	_, err := r.db.NewUpdate().
		Model(cfs).
		Where("id = ?", cfs.ID).
		Exec(ctx)
	return err
}

func (r *CashFlowSummaryRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.NewDelete().
		Model((*schema.CashFlowSummary)(nil)).
		Where("id = ?", id).
		Exec(ctx)
	return err
}
