package usecase

import (
	"context"
	"github.com/Rafiur/wallet_app/internal/domain/entity"
	"github.com/Rafiur/wallet_app/internal/infrastructure/repository"
	"github.com/Rafiur/wallet_app/internal/infrastructure/repository/schema"
)

type ExpenseCategoryService struct {
	ExpenseCategoryRepo repository.ExpenseCategoryRepoInterface
}

func NewExpenseCategoryService(expenseCategoryRepo repository.ExpenseCategoryRepoInterface) *ExpenseCategoryService {
	return &ExpenseCategoryService{
		ExpenseCategoryRepo: expenseCategoryRepo,
	}
}

func (svc *ExpenseCategoryService) Create(ctx context.Context, req *schema.ExpenseCategory) (*schema.ExpenseCategory, error) {
	return svc.ExpenseCategoryRepo.Create(ctx, req)
}

func (svc *ExpenseCategoryService) GetByID(ctx context.Context, id string) (*schema.ExpenseCategory, error) {
	return svc.ExpenseCategoryRepo.GetByID(ctx, id)
}

func (svc *ExpenseCategoryService) GetAll(ctx context.Context, filter *entity.FilterExpenseCategoryListRequest) ([]*schema.ExpenseCategory, error) {
	return svc.ExpenseCategoryRepo.List(ctx, filter)
}

func (svc *ExpenseCategoryService) Update(ctx context.Context, req *schema.ExpenseCategory) (*schema.ExpenseCategory, error) {
	return svc.ExpenseCategoryRepo.Update(ctx, req)
}

func (svc *ExpenseCategoryService) Delete(ctx context.Context, req *entity.CommonDeleteReq) error {
	return svc.ExpenseCategoryRepo.Delete(ctx, req)
}
