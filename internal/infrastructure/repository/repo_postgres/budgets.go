package repo_postgres

import (
	"context"
	"github.com/Rafiur/wallet_app/internal/infrastructure/repository/schema"

	"github.com/uptrace/bun"
)

type BudgetRepository struct {
	db *bun.DB
}

func NewBudgetRepository(db *bun.DB) *BudgetRepository {
	return &BudgetRepository{db: db}
}

func (r *BudgetRepository) Create(ctx context.Context, budget *schema.Budget) error {
	_, err := r.db.NewInsert().Model(budget).Exec(ctx)
	return err
}

func (r *BudgetRepository) GetByID(ctx context.Context, id string) (*schema.Budget, error) {
	budget := new(schema.Budget)
	err := r.db.NewSelect().
		Model(budget).
		Where("id = ?", id).
		Where("deleted_at IS NULL").
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return budget, nil
}

func (r *BudgetRepository) GetByUserID(ctx context.Context, userID string) ([]*schema.Budget, error) {
	var budgets []*schema.Budget
	err := r.db.NewSelect().
		Model(&budgets).
		Where("user_id = ?", userID).
		Where("deleted_at IS NULL").
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return budgets, nil
}

func (r *BudgetRepository) GetByCategoryAndPeriod(ctx context.Context, userID, categoryID, period string) (*schema.Budget, error) {
	budget := new(schema.Budget)
	query := r.db.NewSelect().
		Model(budget).
		Where("user_id = ?", userID).
		Where("period = ?", period).
		Where("deleted_at IS NULL")
	if categoryID != "" {
		query.Where("expense_category_id = ?", categoryID)
	} else {
		query.Where("expense_category_id IS NULL")
	}
	err := query.Scan(ctx)
	if err != nil {
		return nil, err
	}
	return budget, nil
}

func (r *BudgetRepository) Update(ctx context.Context, budget *schema.Budget) error {
	_, err := r.db.NewUpdate().
		Model(budget).
		Where("id = ?", budget.ID).
		Where("deleted_at IS NULL").
		Exec(ctx)
	return err
}

func (r *BudgetRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.NewUpdate().
		Model((*schema.Budget)(nil)).
		Set("deleted_at = CURRENT_TIMESTAMP").
		Where("id = ?", id).
		Where("deleted_at IS NULL").
		Exec(ctx)
	return err
}
