package repo_postgres

import (
	"context"
	"github.com/Rafiur/wallet_app/internal/domain/entity"
	"github.com/Rafiur/wallet_app/internal/infrastructure/repository/schema"
	"github.com/uptrace/bun"
)

type ExpenseCategoryRepo struct {
	db *bun.DB
}

func NewExpenseCategory(db *bun.DB) *ExpenseCategoryRepo {
	return &ExpenseCategoryRepo{db: db}
}

func (repo *ExpenseCategoryRepo) Create(ctx context.Context, req *schema.ExpenseCategory) (*schema.ExpenseCategory, error) {
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

func (repo *ExpenseCategoryRepo) GetByID(ctx context.Context, id string) (*schema.ExpenseCategory, error) {
	var data schema.ExpenseCategory
	err := repo.db.NewSelect().
		Model(&data).
		Where("id = ?", id).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (repo *ExpenseCategoryRepo) List(ctx context.Context, filter *entity.FilterExpenseCategoryListRequest) ([]*schema.ExpenseCategory, error) {
	var data []*schema.ExpenseCategory
	query := repo.db.NewSelect().
		Model(&data).Where("deleted_at IS NULL")
	if filter.ID != "" {
		query = query.Where("id = ?", filter.ID)
	}
	if filter.Name != "" {
		query = query.Where("name = ?", filter.Name)
	}
	if filter.ParentCategoryID != "" {
		query = query.Where("parent_category_id = ?", filter.ParentCategoryID)
	}

	err := query.
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (repo *ExpenseCategoryRepo) Update(ctx context.Context, req *schema.ExpenseCategory) (*schema.ExpenseCategory, error) {
	existing, err := repo.GetByID(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	if req.Name != "" && req.Name != existing.Name {
		existing.Name = req.Name
	}
	if req.ParentCategoryID != nil && (existing.ParentCategoryID == nil || *req.ParentCategoryID != *existing.ParentCategoryID) {
		existing.ParentCategoryID = req.ParentCategoryID
	}

	_, err = repo.db.NewUpdate().
		Model(existing).
		Where("id = ?", req.ID).
		Exec(ctx)

	if err != nil {
		return nil, err
	}
	return req, nil
}

func (repo *ExpenseCategoryRepo) Delete(ctx context.Context, req *entity.CommonDeleteReq) error {

	query := repo.db.NewUpdate().
		Model((*schema.ExpenseCategory)(nil)).
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
