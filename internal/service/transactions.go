package usecase

import (
	"context"
	"errors"
	"github.com/Rafiur/wallet_app/internal/domain/entity"

	"github.com/Rafiur/wallet_app/internal/infrastructure/repository"
	"github.com/Rafiur/wallet_app/internal/infrastructure/repository/schema"
)

type TransactionService struct {
	TransactionRepo repository.TransactionRepoInterface
	AccountRepo     repository.AccountRepoInterface
}

func NewTransactionService(transactionRepo repository.TransactionRepoInterface, accountRepo repository.AccountRepoInterface) *TransactionService {
	return &TransactionService{
		TransactionRepo: transactionRepo,
		AccountRepo:     accountRepo,
	}
}

func (svc *TransactionService) Create(ctx context.Context, req *schema.Transaction) (*schema.Transaction, error) {
	if req.UserID == "" || req.AccountID == "" || req.TransactionName == "" || req.Amount == 0 || req.TransactionType == "" {
		return nil, errors.New("user_id, account_id, transaction_name, amount, and transaction_type are required")
	}
	if req.TransactionType != "income" && req.TransactionType != "expense" {
		return nil, errors.New("transaction_type must be 'income' or 'expense'")
	}
	// Update account balance atomically
	tx, err := svc.TransactionRepo.GetTx(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	account, err := svc.AccountRepo.GetByID(ctx, req.AccountID)
	if err != nil {
		return nil, err
	}
	if req.TransactionType == "expense" && account.Balance+req.Amount < 0 { // Negative amount for expense
		return nil, errors.New("insufficient balance")
	}
	account.Balance += req.Amount
	if _, err := svc.AccountRepo.Update(ctx, account); err != nil {
		return nil, err
	}

	result, err := svc.TransactionRepo.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return result, nil
}

func (svc *TransactionService) GetByID(ctx context.Context, id string) (*schema.Transaction, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return svc.TransactionRepo.GetByID(ctx, id)
}

func (svc *TransactionService) List(ctx context.Context, filter *entity.FilterTransactionListRequest) ([]*schema.Transaction, error) {
	return svc.TransactionRepo.List(ctx, filter)
}

func (svc *TransactionService) Update(ctx context.Context, req *schema.Transaction) (*schema.Transaction, error) {
	if req.ID == "" {
		return nil, errors.New("id is required")
	}
	// Handle balance adjustment
	tx, err := svc.TransactionRepo.GetTx(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	existing, err := svc.TransactionRepo.GetByID(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	if req.Amount != existing.Amount || req.AccountID != existing.AccountID {
		// Revert old amount
		oldAccount, err := svc.AccountRepo.GetByID(ctx, existing.AccountID)
		if err != nil {
			return nil, err
		}
		oldAccount.Balance -= existing.Amount
		if _, err := svc.AccountRepo.Update(ctx, oldAccount); err != nil {
			return nil, err
		}
		// Apply new amount
		newAccount, err := svc.AccountRepo.GetByID(ctx, req.AccountID)
		if err != nil {
			return nil, err
		}
		if req.TransactionType == "expense" && newAccount.Balance+req.Amount < 0 {
			return nil, errors.New("insufficient balance")
		}
		newAccount.Balance += req.Amount
		if _, err := svc.AccountRepo.Update(ctx, newAccount); err != nil {
			return nil, err
		}
	}
	result, err := svc.TransactionRepo.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return result, nil
}

func (svc *TransactionService) Delete(ctx context.Context, req *entity.CommonDeleteReq) error {
	if req.ID == "" && len(req.IDs) == 0 {
		return errors.New("id or ids are required")
	}
	// Revert balances
	tx, err := svc.TransactionRepo.GetTx(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if req.ID != "" {
		txData, err := svc.TransactionRepo.GetByID(ctx, req.ID)
		if err != nil {
			return err
		}
		account, err := svc.AccountRepo.GetByID(ctx, txData.AccountID)
		if err != nil {
			return err
		}
		account.Balance -= txData.Amount
		if _, err := svc.AccountRepo.Update(ctx, account); err != nil {
			return err
		}
	} else {
		txs, err := svc.TransactionRepo.List(ctx, &entity.FilterTransactionListRequest{IDs: req.IDs})
		if err != nil {
			return err
		}
		for _, txData := range txs {
			account, err := svc.AccountRepo.GetByID(ctx, txData.AccountID)
			if err != nil {
				return err
			}
			account.Balance -= txData.Amount
			if _, err := svc.AccountRepo.Update(ctx, account); err != nil {
				return err
			}
		}
	}
	if err := svc.TransactionRepo.Delete(ctx, req); err != nil {
		return err
	}
	return tx.Commit()
}
