package service

import (
	"context"
	"errors"
	"github.com/Rafiur/wallet_app/internal/domain/entity"
	"github.com/Rafiur/wallet_app/internal/infrastructure/repository"
	"github.com/Rafiur/wallet_app/internal/infrastructure/repository/schema"
	"github.com/uptrace/bun"
)

type AccountService struct {
	AccountRepo repository.AccountRepoInterface
}

func NewAccountService(accountRepo repository.AccountRepoInterface) *AccountService {
	return &AccountService{AccountRepo: accountRepo}
}

func (svc *AccountService) GetTx(ctx context.Context) (*bun.Tx, error) {
	return svc.AccountRepo.GetTx(ctx)
}

func (svc *AccountService) Create(ctx context.Context, req *schema.Account) (*schema.Account, error) {
	if req.UserID == "" || req.Name == "" || req.Type == "" || req.Currency == "" {
		return nil, errors.New("user_id, name, type, and currency are required")
	}
	return svc.AccountRepo.Create(ctx, req)
}

func (svc *AccountService) GetByID(ctx context.Context, id string) (*schema.Account, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return svc.AccountRepo.GetByID(ctx, id)
}

func (svc *AccountService) List(ctx context.Context, filter *entity.FilterAccountListRequest) ([]*schema.Account, error) {
	return svc.AccountRepo.List(ctx, filter)
}

func (svc *AccountService) Update(ctx context.Context, req *schema.Account) (*schema.Account, error) {
	if req.ID == "" {
		return nil, errors.New("id is required")
	}
	return svc.AccountRepo.Update(ctx, req)
}

func (svc *AccountService) Delete(ctx context.Context, req *entity.CommonDeleteReq) error {
	if req.ID == "" && len(req.IDs) == 0 {
		return errors.New("id or ids are required")
	}
	return svc.AccountRepo.Delete(ctx, req)
}
