package repository

import (
	"context"
	"github.com/Rafiur/wallet_app/internal/domain/entity"
	"github.com/Rafiur/wallet_app/internal/infrastructure/repository/schema"
)

type ExpenseCategoryRepoInterface interface {
	Create(ctx context.Context, req *schema.ExpenseCategory) (*schema.ExpenseCategory, error)
	GetByID(ctx context.Context, id string) (*schema.ExpenseCategory, error)
	List(ctx context.Context, filter *entity.FilterExpenseCategoryListRequest) ([]*schema.ExpenseCategory, error)
	Update(ctx context.Context, req *schema.ExpenseCategory) (*schema.ExpenseCategory, error)
	Delete(ctx context.Context, req *entity.CommonDeleteReq) error
}
