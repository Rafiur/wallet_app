package usecase

import (
	"context"
	"errors"
	"github.com/Rafiur/wallet_app/internal/domain/entity"
	"github.com/Rafiur/wallet_app/internal/infrastructure/repository"
	"github.com/Rafiur/wallet_app/internal/infrastructure/repository/schema"
)

type ExpenseCategoryService struct {
	ExpenseCategoryRepo repository.ExpenseCategoryRepoInterface
}

func NewExpenseCategoryService(expenseCategoryRepo repository.ExpenseCategoryRepoInterface) *ExpenseCategoryService {
	return &ExpenseCategoryService{ExpenseCategoryRepo: expenseCategoryRepo}
}

func (svc *ExpenseCategoryService) Create(ctx context.Context, req *schema.ExpenseCategory) (*schema.ExpenseCategory, error) {
	if req.Name == "" {
		return nil, errors.New("name is required")
	}
	return svc.ExpenseCategoryRepo.Create(ctx, req)
}

func (svc *ExpenseCategoryService) GetByID(ctx context.Context, id string) (*schema.ExpenseCategory, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return svc.ExpenseCategoryRepo.GetByID(ctx, id)
}

func (svc *ExpenseCategoryService) List(ctx context.Context, filter *entity.FilterExpenseCategoryListRequest) ([]*schema.ExpenseCategory, error) {
	return svc.ExpenseCategoryRepo.List(ctx, filter)
}

func (svc *ExpenseCategoryService) Update(ctx context.Context, req *schema.ExpenseCategory) (*schema.ExpenseCategory, error) {
	if req.ID == "" {
		return nil, errors.New("id is required")
	}
	return svc.ExpenseCategoryRepo.Update(ctx, req)
}

func (svc *ExpenseCategoryService) Delete(ctx context.Context, req *entity.CommonDeleteReq) error {
	if req.ID == "" && len(req.IDs) == 0 {
		return errors.New("id or ids are required")
	}
	return svc.ExpenseCategoryRepo.Delete(ctx, req)
}
