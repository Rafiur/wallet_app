package usecase

import (
	"context"
	"github.com/Rafiur/wallet_app/internal/infrastructure/repository"
	"github.com/Rafiur/wallet_app/internal/infrastructure/repository/schema"
)

type SessionService struct {
	SessionRepoInterface repository.SessionRepoInterface
}

func NewSessionService(sessionRepo repository.SessionRepoInterface) *SessionService {
	return &SessionService{SessionRepoInterface: sessionRepo}
}

func (svc *SessionService) Create(ctx context.Context, req *schema.Session) (*schema.Session, error) {
	return svc.SessionRepoInterface.Create(ctx, req)
}

func (svc *SessionService) GetByID(ctx context.Context, id string) (*schema.Session, error) {
	return svc.SessionRepoInterface.GetByID(ctx, id)
}

func (svc *SessionService) GetByRefreshToken(ctx context.Context, refreshToken string) (*schema.Session, error) {
	return svc.SessionRepoInterface.GetByRefreshToken(ctx, refreshToken)
}

func (svc *SessionService) Delete(ctx context.Context, id string) error {
	return svc.SessionRepoInterface.Delete(ctx, id)
}

func (svc *SessionService) DeleteByUserID(ctx context.Context, userId string) error {
	return svc.SessionRepoInterface.DeleteByUserID(ctx, userId)
}
