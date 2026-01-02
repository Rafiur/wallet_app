package service

import (
	"context"
	"errors"

	"github.com/Rafiur/wallet_app/internal/domain/entity"
	"github.com/Rafiur/wallet_app/internal/infrastructure/repository"
	"github.com/Rafiur/wallet_app/internal/infrastructure/repository/schema"
	"github.com/Rafiur/wallet_app/internal/security"
)

type UserService struct {
	UserRepoInterface repository.UserRepoInterface
	PasswordService   *security.PasswordService
}

func NewUserService(userRepo repository.UserRepoInterface, passwordService *security.PasswordService) *UserService {
	return &UserService{UserRepoInterface: userRepo, PasswordService: passwordService}
}

func (svc *UserService) Create(ctx context.Context, req *schema.User) (*schema.User, error) {
	if req.FullName == "" || req.Email == "" || req.Password == "" {
		return nil, errors.New("full_name, email, and password are required")
	}
	hashedPassword, err := svc.PasswordService.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}
	req.Password = hashedPassword
	return svc.UserRepoInterface.Create(ctx, req)
}

func (svc *UserService) GetByID(ctx context.Context, id string) (*schema.User, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return svc.UserRepoInterface.GetByID(ctx, id)
}

func (svc *UserService) GetByEmail(ctx context.Context, email string) (*schema.User, error) {
	if email == "" {
		return nil, errors.New("email is required")
	}
	return svc.UserRepoInterface.GetByEmail(ctx, email)
}

func (svc *UserService) List(ctx context.Context, filter *entity.FilterUserListRequest) ([]*schema.User, error) {
	return svc.UserRepoInterface.List(ctx, filter)
}

func (svc *UserService) Update(ctx context.Context, req *schema.User) (*schema.User, error) {
	if req.ID == "" {
		return nil, errors.New("id is required")
	}
	return svc.UserRepoInterface.Update(ctx, req)
}

func (svc *UserService) Delete(ctx context.Context, req *entity.CommonDeleteReq) error {
	if req.ID == "" && len(req.IDs) == 0 {
		return errors.New("id or ids are required")
	}
	return svc.UserRepoInterface.Delete(ctx, req)
}
