package repo_postgres

import (
	"context"
	"fmt"
	"github.com/Rafiur/wallet_app/internal/domain/entity"
	"github.com/Rafiur/wallet_app/internal/infrastructure/repository/schema"
	"github.com/uptrace/bun"
	"time"
)

type AccountRepo struct {
	db *bun.DB
}

func NewAccountRepo(db *bun.DB) *AccountRepo {
	return &AccountRepo{db: db}
}

func (repo *AccountRepo) GetTx(ctx context.Context) (*bun.Tx, error) {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	return &tx, err
}

func (repo *AccountRepo) Create(ctx context.Context, req *schema.Account) (*schema.Account, error) {
	_, err := repo.db.NewInsert().
		Model(req).
		ExcludeColumn("deleted_at").
		Returning("*").
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (repo *AccountRepo) GetByID(ctx context.Context, id string) (*schema.Account, error) {
	var data schema.Account
	err := repo.db.NewSelect().
		Model(&data).
		Where("id = ?", id).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (repo *AccountRepo) List(ctx context.Context, filter *entity.FilterAccountListRequest) ([]*schema.Account, error) {
	var data []*schema.Account
	query := repo.db.NewSelect().
		Model(&data).
		Where("deleted_at IS NULL")
	if filter.ID != "" {
		query = query.Where("id = ?", filter.ID)
	}
	if filter.UserID != "" {
		query = query.Where("user_id = ?", filter.UserID)
	}
	if filter.Name != "" {
		query = query.Where("name = ?", filter.Name)
	}

	err := query.
		Order("created_at DESC").
		Scan(ctx)

	if err != nil {
		return nil, err
	}
	return data, nil
}

func (repo *AccountRepo) Update(ctx context.Context, req *schema.Account) (*schema.Account, error) {
	existing, err := repo.GetByID(ctx, req.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve data for update: %w", err)
	}

	if req.Name != "" && req.Name != existing.Name {
		existing.Name = req.Name
	}
	if req.Balance != existing.Balance {
		existing.Balance = req.Balance
	}
	if req.Currency != "" && req.Currency != existing.Currency {
		existing.Currency = req.Currency
	}

	existing.UpdatedAt = time.Now()

	_, err = repo.db.NewUpdate().
		Model(existing).
		ExcludeColumn("deleted_at").
		Where("id = ?", req.ID).
		Exec(ctx)

	if err != nil {
		return nil, err
	}
	return existing, nil
}

func (repo *AccountRepo) Delete(ctx context.Context, req *entity.CommonDeleteReq) error {
	query := repo.db.NewUpdate().
		Model((*schema.Account)(nil)).
		Set("deleted_at = NOW()")
	if len(req.IDs) > 0 {
		query = query.Where("id IN (?)", bun.In(req.IDs))
	} else {
		query = query.Where("id = ?", req.ID)
	}
	_, err := query.Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
