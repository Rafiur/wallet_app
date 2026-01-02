package repo_postgres

import (
	"context"
	"github.com/Rafiur/wallet_app/internal/infrastructure/repository/schema"

	"github.com/uptrace/bun"
)

type CurrencyRepository struct {
	db *bun.DB
}

func NewCurrencyRepository(db *bun.DB) *CurrencyRepository {
	return &CurrencyRepository{db: db}
}

func (r *CurrencyRepository) Create(ctx context.Context, currency *schema.Currency) error {
	_, err := r.db.NewInsert().Model(currency).Exec(ctx)
	return err
}

func (r *CurrencyRepository) GetByCode(ctx context.Context, code string) (*schema.Currency, error) {
	currency := new(schema.Currency)
	err := r.db.NewSelect().
		Model(currency).
		Where("code = ?", code).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return currency, nil
}

func (r *CurrencyRepository) List(ctx context.Context) ([]*schema.Currency, error) {
	var currencies []*schema.Currency
	err := r.db.NewSelect().
		Model(&currencies).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return currencies, nil
}

func (r *CurrencyRepository) Update(ctx context.Context, currency *schema.Currency) error {
	_, err := r.db.NewUpdate().
		Model(currency).
		Where("code = ?", currency.Code).
		Exec(ctx)
	return err
}

func (r *CurrencyRepository) Delete(ctx context.Context, code string) error {
	_, err := r.db.NewDelete().
		Model((*schema.Currency)(nil)).
		Where("code = ?", code).
		Exec(ctx)
	return err
}
