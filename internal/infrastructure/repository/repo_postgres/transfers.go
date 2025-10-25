package repo_postgres

import (
	"context"
	"github.com/Rafiur/wallet_app/internal/domain/entity"
	"github.com/Rafiur/wallet_app/internal/infrastructure/repository/schema"
	"github.com/uptrace/bun"
)

type TransferRepo struct {
	db *bun.DB
}

func NewTransferRepo(db *bun.DB) *TransferRepo {
	return &TransferRepo{db: db}
}

func (repo *TransferRepo) GetTx(ctx context.Context) (*bun.Tx, error) {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	return &tx, err
}

func (repo *TransferRepo) Create(ctx context.Context, req *schema.Transfer) (*schema.Transfer, error) {
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

func (repo *TransferRepo) GetByID(ctx context.Context, id string) (*schema.Transfer, error) {
	var data schema.Transfer
	err := repo.db.NewSelect().
		Model(&data).
		Where("id = ?", id).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (repo *TransferRepo) List(ctx context.Context, filter *entity.FilterTransferListRequest) ([]*schema.Transfer, error) {

	var data []*schema.Transfer
	query := repo.db.NewSelect().Model(&data)

	if len(filter.IDs) > 0 {
		query = query.Where("id IN (?)", bun.In(filter.IDs))
	}
	if filter.FromAccountID != "" {
		query = query.Where("from_account_id = ?", filter.FromAccountID)
	}
	if filter.ToAccountID != "" {
		query = query.Where("to_account_id = ?", filter.ToAccountID)
	}
	if filter.MinAmount > 0 {
		query = query.Where("amount >= ?", filter.MinAmount)
	}
	if filter.MaxAmount > 0 {
		query = query.Where("amount <= ?", filter.MaxAmount)
	}
	if filter.Currency != "" {
		query = query.Where("currency = ?", filter.Currency)
	}
	if filter.Status != "" {
		query = query.Where("status = ?", filter.Status)
	}
	if !filter.StartDate.IsZero() {
		query = query.Where("transfer_date >= ?", filter.StartDate)
	}
	if !filter.EndDate.IsZero() {
		query = query.Where("transfer_date <= ?", filter.EndDate)
	}
	if filter.NoteContains != "" {
		query = query.Where("note ILIKE ?", "%"+filter.NoteContains+"%")
	}
	err := query.Scan(ctx)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (repo *TransferRepo) Update(ctx context.Context, req *schema.Transfer) (*schema.Transfer, error) {
	existing, err := repo.GetByID(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	// Update fields example (add others as needed)
	if req.Amount != 0 {
		existing.Amount = req.Amount
	}
	if req.Note != "" {
		existing.Note = req.Note
	}
	if req.Status != "" {
		existing.Status = req.Status
	}
	_, err = repo.db.NewUpdate().
		Model(existing).
		Where("id = ?", req.ID).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return existing, nil
}

func (repo *TransferRepo) Delete(ctx context.Context, req *entity.CommonDeleteReq) error {
	query := repo.db.NewUpdate().
		Model((*schema.Transfer)(nil)).
		Set("deleted_at = NOW()")

	if len(req.IDs) > 0 {
		query = query.Where("id IN (?)", bun.In(req.IDs))
	}
	if req.ID != "" {
		query = query.Where("id = ?", req.ID)
	}

	_, err := query.Exec(ctx)

	return err
}
