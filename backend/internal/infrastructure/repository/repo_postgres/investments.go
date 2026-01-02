package repo_postgres

import (
	"context"
	"github.com/Rafiur/wallet_app/internal/infrastructure/repository/schema"

	"github.com/uptrace/bun"
)

type InvestmentRepository struct {
	db *bun.DB
}

func NewInvestmentRepository(db *bun.DB) *InvestmentRepository {
	return &InvestmentRepository{db: db}
}

func (r *InvestmentRepository) Create(ctx context.Context, inv *schema.Investment) error {
	_, err := r.db.NewInsert().Model(inv).Exec(ctx)
	return err
}

func (r *InvestmentRepository) GetByID(ctx context.Context, id string) (*schema.Investment, error) {
	inv := new(schema.Investment)
	err := r.db.NewSelect().
		Model(inv).
		Where("id = ?", id).
		Where("deleted_at IS NULL").
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return inv, nil
}

func (r *InvestmentRepository) GetByUserID(ctx context.Context, userID string) ([]*schema.Investment, error) {
	var investments []*schema.Investment
	err := r.db.NewSelect().
		Model(&investments).
		Where("user_id = ?", userID).
		Where("deleted_at IS NULL").
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return investments, nil
}

func (r *InvestmentRepository) Update(ctx context.Context, inv *schema.Investment) error {
	_, err := r.db.NewUpdate().
		Model(inv).
		Where("id = ?", inv.ID).
		Where("deleted_at IS NULL").
		Exec(ctx)
	return err
}

func (r *InvestmentRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.NewUpdate().
		Model((*schema.Investment)(nil)).
		Set("deleted_at = CURRENT_TIMESTAMP").
		Where("id = ?", id).
		Where("deleted_at IS NULL").
		Exec(ctx)
	return err
}
