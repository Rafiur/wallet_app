package repo_postgres

import (
	"context"
	"github.com/Rafiur/wallet_app/internal/infrastructure/repository/schema"

	"github.com/uptrace/bun"
)

type BankRepository struct {
	db *bun.DB
}

func NewBankRepository(db *bun.DB) *BankRepository {
	return &BankRepository{db: db}
}

func (r *BankRepository) Create(ctx context.Context, bank *schema.Bank) error {
	_, err := r.db.NewInsert().Model(bank).Exec(ctx)
	return err
}

func (r *BankRepository) GetByID(ctx context.Context, id string) (*schema.Bank, error) {
	bank := new(schema.Bank)
	err := r.db.NewSelect().
		Model(bank).
		Relation("Account").
		Where("b.id = ?", id).
		Where("b.deleted_at IS NULL").
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return bank, nil
}

func (r *BankRepository) GetByUserID(ctx context.Context, userID string) ([]*schema.Bank, error) {
	var banks []*schema.Bank
	err := r.db.NewSelect().
		Model(&banks).
		Relation("Account").
		Where("b.user_id = ?", userID).
		Where("b.deleted_at IS NULL").
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return banks, nil
}

func (r *BankRepository) GetByAccountID(ctx context.Context, accountID string) (*schema.Bank, error) {
	bank := new(schema.Bank)
	err := r.db.NewSelect().
		Model(bank).
		Relation("Account").
		Where("b.account_id = ?", accountID).
		Where("b.deleted_at IS NULL").
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return bank, nil
}

func (r *BankRepository) Update(ctx context.Context, bank *schema.Bank) error {
	_, err := r.db.NewUpdate().
		Model(bank).
		Where("id = ?", bank.ID).
		Where("deleted_at IS NULL").
		Exec(ctx)
	return err
}

func (r *BankRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.NewUpdate().
		Model((*schema.Bank)(nil)).
		Set("deleted_at = CURRENT_TIMESTAMP").
		Where("id = ?", id).
		Where("deleted_at IS NULL").
		Exec(ctx)
	return err
}
