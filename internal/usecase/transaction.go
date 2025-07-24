package usecase

import (
	"context"
	"github.com/Rafiur/wallet_app/internal/domain/entity"
	"github.com/Rafiur/wallet_app/internal/infrastructure/repository"
	"github.com/Rafiur/wallet_app/internal/infrastructure/repository/schema"
)

type TransactionService struct {
	TransactionRepoInterface repository.TransactionRepoInterface
}

func NewTransactionService(transactionRepo repository.TransactionRepoInterface) *TransactionService {
	return &TransactionService{TransactionRepoInterface: transactionRepo}
}

func (svc *TransactionService) Create(ctx context.Context, req *schema.Transaction) (*schema.Transaction, error) {
	return svc.TransactionRepoInterface.Create(ctx, req)
}

func (svc *TransactionService) GetByID(ctx context.Context, id string) (*schema.Transaction, error) {
	return svc.TransactionRepoInterface.GetByID(ctx, id)
}

func (svc *TransactionService) List(ctx context.Context, filter *entity.FilterTransactionListRequest) ([]*schema.Transaction, error) {
	return svc.TransactionRepoInterface.List(ctx, filter)
}

func (svc *TransactionService) Update(ctx context.Context, req *schema.Transaction) (*schema.Transaction, error) {
	return svc.TransactionRepoInterface.Update(ctx, req)
}

func (svc *TransactionService) Delete(ctx context.Context, req *entity.CommonDeleteReq) error {
	return svc.TransactionRepoInterface.Delete(ctx, req)
}
