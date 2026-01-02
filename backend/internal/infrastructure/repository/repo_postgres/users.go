package repo_postgres

import (
	"context"

	"github.com/Rafiur/wallet_app/internal/domain/entity"
	"github.com/Rafiur/wallet_app/internal/infrastructure/repository/schema"
	"github.com/uptrace/bun"
)

type UserRepo struct {
	db *bun.DB
}

func NewUserRepo(db *bun.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (repo *UserRepo) Create(ctx context.Context, req *schema.User) (*schema.User, error) {
	_, err := repo.db.NewInsert().
		Model(req).
		ExcludeColumn("deleted_at").
		Exec(ctx)

	if err != nil {
		return nil, err
	}
	return req, nil
}

func (repo *UserRepo) GetByID(ctx context.Context, id string) (*schema.User, error) {
	var data schema.User
	err := repo.db.NewSelect().
		Model(&data).
		Where("id = ?", id).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (repo *UserRepo) GetByEmail(ctx context.Context, email string) (*schema.User, error) {
	var data schema.User
	err := repo.db.NewSelect().
		Model(&data).
		Where("email = ?", email).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (repo *UserRepo) List(ctx context.Context, filter *entity.FilterUserListRequest) ([]*schema.User, error) {
	var data []*schema.User
	query := repo.db.NewSelect().
		Model(&data).
		Where("deleted_at IS NULL")
	if filter.ID != "" {
		query = query.Where("id = ?", filter.ID)
	}
	if filter.FullName != "" {
		query = query.Where("full_name LIKE ?", "%"+filter.FullName+"%")
	}
	err := query.Scan(ctx)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (repo *UserRepo) Update(ctx context.Context, req *schema.User) (*schema.User, error) {
	existing, err := repo.GetByID(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	if req.FullName != "" {
		existing.FullName = req.FullName
	}

	_, err = repo.db.NewUpdate().
		Model(existing).
		Exec(ctx)

	if err != nil {
		return nil, err
	}
	return existing, nil
}

func (repo *UserRepo) Delete(ctx context.Context, req *entity.CommonDeleteReq) error {
	query := repo.db.NewUpdate().
		Model((*schema.User)(nil)).
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
