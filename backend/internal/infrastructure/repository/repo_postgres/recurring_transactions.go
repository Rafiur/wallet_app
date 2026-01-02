package repo_postgres

import (
	"context"
	"github.com/Rafiur/wallet_app/internal/infrastructure/repository/schema"

	"github.com/uptrace/bun"
	"time"
)

type RecurringTransactionRepository struct {
	db *bun.DB
}

func NewRecurringTransactionRepository(db *bun.DB) *RecurringTransactionRepository {
	return &RecurringTransactionRepository{db: db}
}

func (r *RecurringTransactionRepository) Create(ctx context.Context, rt *schema.RecurringTransaction) error {
	_, err := r.db.NewInsert().Model(rt).Exec(ctx)
	return err
}

func (r *RecurringTransactionRepository) GetByID(ctx context.Context, id string) (*schema.RecurringTransaction, error) {
	rt := new(schema.RecurringTransaction)
	err := r.db.NewSelect().
		Model(rt).
		Where("id = ?", id).
		Where("deleted_at IS NULL").
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return rt, nil
}

func (r *RecurringTransactionRepository) GetByUserID(ctx context.Context, userID string) ([]*schema.RecurringTransaction, error) {
	var rts []*schema.RecurringTransaction
	err := r.db.NewSelect().
		Model(&rts).
		Where("user_id = ?", userID).
		Where("deleted_at IS NULL").
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return rts, nil
}

func (r *RecurringTransactionRepository) GetDueByDate(ctx context.Context, dueDate time.Time) ([]*schema.RecurringTransaction, error) {
	var rts []*schema.RecurringTransaction
	err := r.db.NewSelect().
		Model(&rts).
		Where("next_due_date <= ?", dueDate).
		Where("is_active = ?", true).
		Where("deleted_at IS NULL").
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return rts, nil
}

func (r *RecurringTransactionRepository) Update(ctx context.Context, rt *schema.RecurringTransaction) error {
	_, err := r.db.NewUpdate().
		Model(rt).
		Where("id = ?", rt.ID).
		Where("deleted_at IS NULL").
		Exec(ctx)
	return err
}

func (r *RecurringTransactionRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.NewUpdate().
		Model((*schema.RecurringTransaction)(nil)).
		Set("deleted_at = CURRENT_TIMESTAMP").
		Where("id = ?", id).
		Where("deleted_at IS NULL").
		Exec(ctx)
	return err
}
