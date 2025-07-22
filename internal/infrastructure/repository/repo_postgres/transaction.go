package repo_postgres

import (
	"context"

	"github.com/Rafiur/wallet_app/internal/domain/entity"
	"github.com/Rafiur/wallet_app/internal/infrastructure/repository/schema"
	"github.com/uptrace/bun"
)

type TransactionRepo struct {
	db *bun.DB
}

func NewTransactionRepo(db *bun.DB) *TransactionRepo {
	return &TransactionRepo{db: db}
}

func (repo *TransactionRepo) GetTx(ctx context.Context) (*bun.Tx, error) {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	return &tx, err
}

func (repo *TransactionRepo) Create(ctx context.Context, req *schema.Transaction) (*schema.Transaction, error) {
	_, err := repo.db.NewInsert().
		Model(req).
		ExcludeColumn("deleted_at").
		Returning("*").
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return req, err
}

func (repo *TransactionRepo) GetByID(ctx context.Context, id string) (*schema.Transaction, error) {
	var data schema.Transaction
	err := repo.db.NewSelect().
		Model(&data).
		Where("id = ?", id).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return &data, err
}

func (repo *TransactionRepo) List(ctx context.Context, filter *entity.FilterTransactionListRequest) ([]*schema.Transaction, error) {
	var data []*schema.Transaction

	query := repo.db.NewSelect().Model(&data).Where("deleted_at IS NULL")

	if filter.ID != "" {
		query = query.Where("id = ?", filter.ID)
	}
	if filter.UserID != "" {
		query = query.Where("user_id = ?", filter.UserID)
	}
	if filter.AccountID != "" {
		query = query.Where("account_id = ?", filter.AccountID)
	}
	if filter.ExpenseCategoryID != "" {
		query = query.Where("expense_category_id = ?", filter.ExpenseCategoryID)
	}
	if filter.TransactionType != "" {
		query = query.Where("transaction_type = ?", filter.TransactionType)
	}
	if !filter.StartDate.IsZero() {
		query = query.Where("transaction_date >= ?", filter.StartDate)
	}
	if !filter.EndDate.IsZero() {
		query = query.Where("transaction_date <= ?", filter.EndDate)
	}
	if filter.SearchBy != "" {
		query = query.Where("transaction_name ILIKE ? OR note ILIKE ?", "%"+filter.SearchBy+"%", "%"+filter.SearchBy+"%")
	}
	if len(filter.Tags) > 0 {
		query = query.Where("tags && ?", filter.Tags)
	}

	err := query.Scan(ctx)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (repo *TransactionRepo) Update(ctx context.Context, req *schema.Transaction) (*schema.Transaction, error) {
	existing, err := repo.GetByID(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	// Update only non-empty or changed fields
	if req.TransactionName != "" && req.TransactionName != existing.TransactionName {
		existing.TransactionName = req.TransactionName
	}
	if !req.TransactionDate.IsZero() && req.TransactionDate != existing.TransactionDate {
		existing.TransactionDate = req.TransactionDate
	}
	if req.Amount != existing.Amount {
		existing.Amount = req.Amount
	}
	if req.AccountID != "" && req.AccountID != existing.AccountID {
		existing.AccountID = req.AccountID
	}
	if req.ExpenseCategoryID != "" && req.ExpenseCategoryID != existing.ExpenseCategoryID {
		existing.ExpenseCategoryID = req.ExpenseCategoryID
	}
	if req.Note != "" && req.Note != existing.Note {
		existing.Note = req.Note
	}
	if len(req.Tags) > 0 {
		existing.Tags = req.Tags
	}
	if req.TransactionType != "" && req.TransactionType != existing.TransactionType {
		existing.TransactionType = req.TransactionType
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

func (repo *TransactionRepo) Delete(ctx context.Context, req *entity.CommonDeleteReq) error {
	query := repo.db.NewUpdate().
		Model((*schema.Transaction)(nil)).
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
