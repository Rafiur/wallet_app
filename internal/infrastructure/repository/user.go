package repository

import (
	"context"
	"github.com/Rafiur/wallet_app/internal/domain/entity"
	"github.com/Rafiur/wallet_app/internal/infrastructure/repository/schema"
)

type UserRepoInterface interface {
	Create(ctx context.Context, req *schema.User) (*schema.User, error)
	GetByID(ctx context.Context, id string) (*schema.User, error)
	List(ctx context.Context, filter *entity.FilterUserListRequest) ([]*schema.User, error)
	Update(ctx context.Context, req *schema.User) (*schema.User, error)
	Delete(ctx context.Context, req *entity.CommonDeleteReq) error
}
