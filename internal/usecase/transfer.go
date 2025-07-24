package usecase

import (
	"context"
	"github.com/Rafiur/wallet_app/internal/domain/entity"
	"github.com/Rafiur/wallet_app/internal/infrastructure/repository"
	"github.com/Rafiur/wallet_app/internal/infrastructure/repository/schema"
)

type TransferService struct {
	TransferRepoInterface repository.TransferRepoInterface
}

func NewTransferService(transferRepo repository.TransferRepoInterface) *TransferService {
	return &TransferService{TransferRepoInterface: transferRepo}
}

func (svc *TransferService) Create(ctx context.Context, req *schema.Transfer) (*schema.Transfer, error) {
	return svc.TransferRepoInterface.Create(ctx, req)
}

func (svc *TransferService) GetByID(ctx context.Context, id string) (*schema.Transfer, error) {
	return svc.TransferRepoInterface.GetByID(ctx, id)
}

func (svc *TransferService) List(ctx context.Context, filter *entity.FilterTransferListRequest) ([]*schema.Transfer, error) {
	return svc.TransferRepoInterface.List(ctx, filter)
}

func (svc *TransferService) Update(ctx context.Context, req *schema.Transfer) (*schema.Transfer, error) {
	return svc.TransferRepoInterface.Update(ctx, req)
}

func (svc *TransferService) Delete(ctx context.Context, req *entity.CommonDeleteReq) error {
	return svc.TransferRepoInterface.Delete(ctx, req)
}
