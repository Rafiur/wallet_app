package repo_postgres

import (
	"context"
	"github.com/Rafiur/wallet_app/internal/infrastructure/repository/schema"

	"github.com/uptrace/bun"
)

type AccountCurrenciesRepository struct {
	db *bun.DB
}

func NewAccountCurrenciesRepository(db *bun.DB) *AccountCurrenciesRepository {
	return &AccountCurrenciesRepository{db: db}
}

func (r *AccountCurrenciesRepository) Create(ctx context.Context, ac *schema.AccountCurrencies) error {
	_, err := r.db.NewInsert().Model(ac).Exec(ctx)
	return err
}

func (r *AccountCurrenciesRepository) GetByID(ctx context.Context, id string) (*schema.AccountCurrencies, error) {
	ac := new(schema.AccountCurrencies)
	err := r.db.NewSelect().
		Model(ac).
		Relation("Account").
		Where("ac.id = ?", id).
		Where("ac.deleted_at IS NULL").
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return ac, nil
}

func (r *AccountCurrenciesRepository) GetByAccountID(ctx context.Context, accountID string) ([]*schema.AccountCurrencies, error) {
	var acs []*schema.AccountCurrencies
	err := r.db.NewSelect().
		Model(&acs).
		Relation("Account").
		Where("ac.account_id = ?", accountID).
		Where("ac.deleted_at IS NULL").
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return acs, nil
}

func (r *AccountCurrenciesRepository) Update(ctx context.Context, ac *schema.AccountCurrencies) error {
	_, err := r.db.NewUpdate().
		Model(ac).
		Where("id = ?", ac.ID).
		Where("deleted_at IS NULL").
		Exec(ctx)
	return err
}

func (r *AccountCurrenciesRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.NewUpdate().
		Model((*schema.AccountCurrencies)(nil)).
		Set("deleted_at = CURRENT_TIMESTAMP").
		Where("id = ?", id).
		Where("deleted_at IS NULL").
		Exec(ctx)
	return err
}
