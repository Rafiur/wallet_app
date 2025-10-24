package repository

import (
	"context"
	"github.com/Rafiur/wallet_app/internal/domain/entity"
	"github.com/Rafiur/wallet_app/internal/infrastructure/repository/schema"
	"github.com/uptrace/bun"
)

type TransactionRepoInterface interface {
	GetTx(ctx context.Context) (*bun.Tx, error)
	Create(ctx context.Context, req *schema.Transaction) (*schema.Transaction, error)
	GetByID(ctx context.Context, id string) (*schema.Transaction, error)
	List(ctx context.Context, filter *entity.FilterTransactionListRequest) ([]*schema.Transaction, error)
	Update(ctx context.Context, req *schema.Transaction) (*schema.Transaction, error)
	Delete(ctx context.Context, req *entity.CommonDeleteReq) error
}
