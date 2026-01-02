package service

import (
	"context"
	"errors"
	"github.com/Rafiur/wallet_app/internal/infrastructure/repository"
	"github.com/Rafiur/wallet_app/internal/infrastructure/repository/schema"

	"time"
)

type RecurringTransactionService struct {
	repo repository.RecurringTransactionRepository
}

func NewRecurringTransactionService(repo repository.RecurringTransactionRepository) *RecurringTransactionService {
	return &RecurringTransactionService{repo: repo}
}

func (svc *RecurringTransactionService) Create(ctx context.Context, rt *schema.RecurringTransaction) error {
	if rt.UserID == "" || rt.AccountID == "" || rt.Name == "" || rt.Frequency == "" || rt.Amount == 0 {
		return errors.New("user_id, account_id, name, frequency, and amount are required")
	}
	return svc.repo.Create(ctx, rt)
}

func (svc *RecurringTransactionService) GetByID(ctx context.Context, id string) (*schema.RecurringTransaction, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return svc.repo.GetByID(ctx, id)
}

func (svc *RecurringTransactionService) GetByUserID(ctx context.Context, userID string) ([]*schema.RecurringTransaction, error) {
	if userID == "" {
		return nil, errors.New("user_id is required")
	}
	return svc.repo.GetByUserID(ctx, userID)
}

func (svc *RecurringTransactionService) GetDueByDate(ctx context.Context, dueDate time.Time) ([]*schema.RecurringTransaction, error) {
	return svc.repo.GetDueByDate(ctx, dueDate)
}

func (svc *RecurringTransactionService) Update(ctx context.Context, rt *schema.RecurringTransaction) error {
	if rt.ID == "" {
		return errors.New("id is required")
	}
	return svc.repo.Update(ctx, rt)
}

func (svc *RecurringTransactionService) Delete(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("id is required")
	}
	return svc.repo.Delete(ctx, id)
}
