package service

import (
	"context"
	"errors"
	"github.com/Rafiur/wallet_app/internal/domain/entity"
	"github.com/Rafiur/wallet_app/internal/infrastructure/repository"
	"github.com/Rafiur/wallet_app/internal/infrastructure/repository/schema"
)

type TransferService struct {
	TransferRepo repository.TransferRepoInterface
	AccountRepo  repository.AccountRepoInterface
}

func NewTransferService(transferRepo repository.TransferRepoInterface, accountRepo repository.AccountRepoInterface) *TransferService {
	return &TransferService{
		TransferRepo: transferRepo,
		AccountRepo:  accountRepo,
	}
}

func (svc *TransferService) Create(ctx context.Context, req *schema.Transfer) (*schema.Transfer, error) {
	if req.FromAccountID == "" || req.ToAccountID == "" || req.Amount <= 0 || req.Currency == "" {
		return nil, errors.New("from_account_id, to_account_id, amount, and currency are required")
	}
	if req.FromAccountID == req.ToAccountID {
		return nil, errors.New("cannot transfer to the same account")
	}
	tx, err := svc.TransferRepo.GetTx(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	accountRepo := svc.AccountRepo.WithTx(tx)
	transferRepo := svc.TransferRepo.WithTx(tx)

	fromAccount, err := accountRepo.GetByID(ctx, req.FromAccountID)
	if err != nil {
		return nil, err
	}
	if fromAccount.Balance < req.Amount {
		return nil, errors.New("insufficient balance in from_account")
	}
	fromAccount.Balance -= req.Amount
	if _, err := accountRepo.Update(ctx, fromAccount); err != nil {
		return nil, err
	}

	toAccount, err := accountRepo.GetByID(ctx, req.ToAccountID)
	if err != nil {
		return nil, err
	}
	if req.ExchangeRate != nil {
		toAccount.Balance += req.Amount * *req.ExchangeRate
	} else {
		toAccount.Balance += req.Amount
	}
	if _, err := accountRepo.Update(ctx, toAccount); err != nil {
		return nil, err
	}

	req.Status = "completed"
	result, err := transferRepo.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return result, nil
}

func (svc *TransferService) GetByID(ctx context.Context, id string) (*schema.Transfer, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return svc.TransferRepo.GetByID(ctx, id)
}

func (svc *TransferService) List(ctx context.Context, filter *entity.FilterTransferListRequest) ([]*schema.Transfer, error) {
	return svc.TransferRepo.List(ctx, filter)
}

func (svc *TransferService) Update(ctx context.Context, req *schema.Transfer) (*schema.Transfer, error) {
	if req.ID == "" {
		return nil, errors.New("id is required")
	}
	tx, err := svc.TransferRepo.GetTx(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	accountRepo := svc.AccountRepo.WithTx(tx)
	transferRepo := svc.TransferRepo.WithTx(tx)

	existing, err := transferRepo.GetByID(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	if req.Amount != existing.Amount || req.FromAccountID != existing.FromAccountID || req.ToAccountID != existing.ToAccountID {
		// Revert old transfer
		fromAccount, err := accountRepo.GetByID(ctx, existing.FromAccountID)
		if err != nil {
			return nil, err
		}
		fromAccount.Balance += existing.Amount
		if _, err := accountRepo.Update(ctx, fromAccount); err != nil {
			return nil, err
		}
		toAccount, err := accountRepo.GetByID(ctx, existing.ToAccountID)
		if err != nil {
			return nil, err
		}
		if existing.ExchangeRate != nil {
			toAccount.Balance -= existing.Amount * *existing.ExchangeRate
		} else {
			toAccount.Balance -= existing.Amount
		}
		if _, err := accountRepo.Update(ctx, toAccount); err != nil {
			return nil, err
		}
		// Apply new transfer
		newFromAccount, err := accountRepo.GetByID(ctx, req.FromAccountID)
		if err != nil {
			return nil, err
		}
		if newFromAccount.Balance < req.Amount {
			return nil, errors.New("insufficient balance in from_account")
		}
		newFromAccount.Balance -= req.Amount
		if _, err := accountRepo.Update(ctx, newFromAccount); err != nil {
			return nil, err
		}
		newToAccount, err := accountRepo.GetByID(ctx, req.ToAccountID)
		if err != nil {
			return nil, err
		}
		if req.ExchangeRate != nil {
			newToAccount.Balance += req.Amount * *req.ExchangeRate
		} else {
			newToAccount.Balance += req.Amount
		}
		if _, err := accountRepo.Update(ctx, newToAccount); err != nil {
			return nil, err
		}
	}
	result, err := transferRepo.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return result, nil
}

func (svc *TransferService) Delete(ctx context.Context, req *entity.CommonDeleteReq) error {
	if req.ID == "" && len(req.IDs) == 0 {
		return errors.New("id or ids are required")
	}
	tx, err := svc.TransferRepo.GetTx(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	accountRepo := svc.AccountRepo.WithTx(tx)
	transferRepo := svc.TransferRepo.WithTx(tx)

	if req.ID != "" {
		transfer, err := transferRepo.GetByID(ctx, req.ID)
		if err != nil {
			return err
		}
		fromAccount, err := accountRepo.GetByID(ctx, transfer.FromAccountID)
		if err != nil {
			return err
		}
		fromAccount.Balance += transfer.Amount
		if _, err := accountRepo.Update(ctx, fromAccount); err != nil {
			return err
		}
		toAccount, err := accountRepo.GetByID(ctx, transfer.ToAccountID)
		if err != nil {
			return err
		}
		if transfer.ExchangeRate != nil {
			toAccount.Balance -= transfer.Amount * *transfer.ExchangeRate
		} else {
			toAccount.Balance -= transfer.Amount
		}
		if _, err := accountRepo.Update(ctx, toAccount); err != nil {
			return err
		}
	} else {
		transfers, err := transferRepo.List(ctx, &entity.FilterTransferListRequest{IDs: req.IDs})
		if err != nil {
			return err
		}
		for _, transfer := range transfers {
			fromAccount, err := accountRepo.GetByID(ctx, transfer.FromAccountID)
			if err != nil {
				return err
			}
			fromAccount.Balance += transfer.Amount
			if _, err := accountRepo.Update(ctx, fromAccount); err != nil {
				return err
			}
			toAccount, err := accountRepo.GetByID(ctx, transfer.ToAccountID)
			if err != nil {
				return err
			}
			if transfer.ExchangeRate != nil {
				toAccount.Balance -= transfer.Amount * *transfer.ExchangeRate
			} else {
				toAccount.Balance -= transfer.Amount
			}
			if _, err := accountRepo.Update(ctx, toAccount); err != nil {
				return err
			}
		}
	}
	if err := transferRepo.Delete(ctx, req); err != nil {
		return err
	}
	return tx.Commit()
}
