package service

import (
	"context"
	"errors"
	"github.com/Rafiur/wallet_app/internal/infrastructure/repository"
	"github.com/Rafiur/wallet_app/internal/infrastructure/repository/schema"
)

type BankService struct {
	repo repository.BankRepository
}

func NewBankService(repo repository.BankRepository) *BankService {
	return &BankService{repo: repo}
}

func (svc *BankService) Create(ctx context.Context, bank *schema.Bank) error {
	if bank.Name == "" || bank.AccountID == "" || bank.AccountNumber == "" {
		return errors.New("name, account_id, and account_number are required")
	}
	return svc.repo.Create(ctx, bank)
}

func (svc *BankService) GetByID(ctx context.Context, id string) (*schema.Bank, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return svc.repo.GetByID(ctx, id)
}

func (svc *BankService) GetByUserID(ctx context.Context, userID string) ([]*schema.Bank, error) {
	if userID == "" {
		return nil, errors.New("user_id is required")
	}
	return svc.repo.GetByUserID(ctx, userID)
}

func (svc *BankService) GetByAccountID(ctx context.Context, accountID string) (*schema.Bank, error) {
	if accountID == "" {
		return nil, errors.New("account_id is required")
	}
	return svc.repo.GetByAccountID(ctx, accountID)
}

func (svc *BankService) Update(ctx context.Context, bank *schema.Bank) error {
	if bank.ID == "" {
		return errors.New("id is required")
	}
	return svc.repo.Update(ctx, bank)
}

func (svc *BankService) Delete(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("id is required")
	}
	return svc.repo.Delete(ctx, id)
}
