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

	fromAccount, err := svc.AccountRepo.GetByID(ctx, req.FromAccountID)
	if err != nil {
		return nil, err
	}
	if fromAccount.Balance < req.Amount {
		return nil, errors.New("insufficient balance in from_account")
	}
	fromAccount.Balance -= req.Amount
	if _, err := svc.AccountRepo.Update(ctx, fromAccount); err != nil {
		return nil, err
	}

	toAccount, err := svc.AccountRepo.GetByID(ctx, req.ToAccountID)
	if err != nil {
		return nil, err
	}
	if req.ExchangeRate != nil {
		toAccount.Balance += req.Amount * *req.ExchangeRate
	} else {
		toAccount.Balance += req.Amount
	}
	if _, err := svc.AccountRepo.Update(ctx, toAccount); err != nil {
		return nil, err
	}

	req.Status = "completed"
	result, err := svc.TransferRepo.Create(ctx, req)
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

	existing, err := svc.TransferRepo.GetByID(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	if req.Amount != existing.Amount || req.FromAccountID != existing.FromAccountID || req.ToAccountID != existing.ToAccountID {
		// Revert old transfer
		fromAccount, err := svc.AccountRepo.GetByID(ctx, existing.FromAccountID)
		if err != nil {
			return nil, err
		}
		fromAccount.Balance += existing.Amount
		if _, err := svc.AccountRepo.Update(ctx, fromAccount); err != nil {
			return nil, err
		}
		toAccount, err := svc.AccountRepo.GetByID(ctx, existing.ToAccountID)
		if err != nil {
			return nil, err
		}
		if existing.ExchangeRate != nil {
			toAccount.Balance -= existing.Amount * *existing.ExchangeRate
		} else {
			toAccount.Balance -= existing.Amount
		}
		if _, err := svc.AccountRepo.Update(ctx, toAccount); err != nil {
			return nil, err
		}
		// Apply new transfer
		newFromAccount, err := svc.AccountRepo.GetByID(ctx, req.FromAccountID)
		if err != nil {
			return nil, err
		}
		if newFromAccount.Balance < req.Amount {
			return nil, errors.New("insufficient balance in from_account")
		}
		newFromAccount.Balance -= req.Amount
		if _, err := svc.AccountRepo.Update(ctx, newFromAccount); err != nil {
			return nil, err
		}
		newToAccount, err := svc.AccountRepo.GetByID(ctx, req.ToAccountID)
		if err != nil {
			return nil, err
		}
		if req.ExchangeRate != nil {
			newToAccount.Balance += req.Amount * *req.ExchangeRate
		} else {
			newToAccount.Balance += req.Amount
		}
		if _, err := svc.AccountRepo.Update(ctx, newToAccount); err != nil {
			return nil, err
		}
	}
	result, err := svc.TransferRepo.Update(ctx, req)
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

	if req.ID != "" {
		transfer, err := svc.TransferRepo.GetByID(ctx, req.ID)
		if err != nil {
			return err
		}
		fromAccount, err := svc.AccountRepo.GetByID(ctx, transfer.FromAccountID)
		if err != nil {
			return err
		}
		fromAccount.Balance += transfer.Amount
		if _, err := svc.AccountRepo.Update(ctx, fromAccount); err != nil {
			return err
		}
		toAccount, err := svc.AccountRepo.GetByID(ctx, transfer.ToAccountID)
		if err != nil {
			return err
		}
		if transfer.ExchangeRate != nil {
			toAccount.Balance -= transfer.Amount * *transfer.ExchangeRate
		} else {
			toAccount.Balance -= transfer.Amount
		}
		if _, err := svc.AccountRepo.Update(ctx, toAccount); err != nil {
			return err
		}
	} else {
		transfers, err := svc.TransferRepo.List(ctx, &entity.FilterTransferListRequest{IDs: req.IDs})
		if err != nil {
			return err
		}
		for _, transfer := range transfers {
			fromAccount, err := svc.AccountRepo.GetByID(ctx, transfer.FromAccountID)
			if err != nil {
				return err
			}
			fromAccount.Balance += transfer.Amount
			if _, err := svc.AccountRepo.Update(ctx, fromAccount); err != nil {
				return err
			}
			toAccount, err := svc.AccountRepo.GetByID(ctx, transfer.ToAccountID)
			if err != nil {
				return err
			}
			if transfer.ExchangeRate != nil {
				toAccount.Balance -= transfer.Amount * *transfer.ExchangeRate
			} else {
				toAccount.Balance -= transfer.Amount
			}
			if _, err := svc.AccountRepo.Update(ctx, toAccount); err != nil {
				return err
			}
		}
	}
	if err := svc.TransferRepo.Delete(ctx, req); err != nil {
		return err
	}
	return tx.Commit()
}
