package repository

import (
	"context"
	"github.com/Rafiur/wallet_app/internal/domain/entity"
	"github.com/Rafiur/wallet_app/internal/infrastructure/repository/schema"
	"github.com/uptrace/bun"
)

type AccountRepoInterface interface {
	GetTx(ctx context.Context) (*bun.Tx, error)
	// WithTx returns a repo whose operations run against the given transaction
	// instead of the default connection, so a caller can compose several writes
	// (e.g. balance update + transaction insert) into one atomic unit.
	WithTx(tx bun.IDB) AccountRepoInterface

	Create(ctx context.Context, req *schema.Account) (*schema.Account, error)
	GetByID(ctx context.Context, id string) (*schema.Account, error)
	List(ctx context.Context, filter *entity.FilterAccountListRequest) ([]*schema.Account, error)
	Update(ctx context.Context, req *schema.Account) (*schema.Account, error)
	Delete(ctx context.Context, req *entity.CommonDeleteReq) error
}
