package repository

import (
	"context"
	"github.com/Rafiur/wallet_app/internal/domain/entity"
	"github.com/Rafiur/wallet_app/internal/infrastructure/repository/schema"
	"github.com/uptrace/bun"
)

type TransferRepoInterface interface {
	GetTx(ctx context.Context) (*bun.Tx, error)
	Create(ctx context.Context, req *schema.Transfer) (*schema.Transfer, error)
	GetByID(ctx context.Context, id string) (*schema.Transfer, error)
	List(ctx context.Context, filter *entity.FilterTransferListRequest) ([]*schema.Transfer, error)
	Update(ctx context.Context, req *schema.Transfer) (*schema.Transfer, error)
	Delete(ctx context.Context, req *entity.CommonDeleteReq) error
}
