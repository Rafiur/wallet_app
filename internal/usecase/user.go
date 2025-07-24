package usecase

import (
	"context"
	"github.com/Rafiur/wallet_app/internal/domain/entity"
	"github.com/Rafiur/wallet_app/internal/infrastructure/repository"
	"github.com/Rafiur/wallet_app/internal/infrastructure/repository/schema"
)

type UserService struct {
	UserRepoInterface repository.UserRepoInterface
}

func NewUserService(userRepo repository.UserRepoInterface) *UserService {
	return &UserService{UserRepoInterface: userRepo}
}

func (svc *UserService) Create(ctx context.Context, req *schema.User) (*schema.User, error) {
	return svc.UserRepoInterface.Create(ctx, req)
}

func (svc *UserService) GetByID(ctx context.Context, id string) (*schema.User, error) {
	return svc.UserRepoInterface.GetByID(ctx, id)
}

func (svc *UserService) List(ctx context.Context, filter *entity.FilterUserListRequest) ([]*schema.User, error) {
	return svc.UserRepoInterface.List(ctx, filter)
}

func (svc *UserService) Update(ctx context.Context, req *schema.User) (*schema.User, error) {
	return svc.UserRepoInterface.Update(ctx, req)
}

func (svc *UserService) Delete(ctx context.Context, req *entity.CommonDeleteReq) error {
	return svc.UserRepoInterface.Delete(ctx, req)
}
