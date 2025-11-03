package service

import (
	"context"
	"errors"

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
	if req.UserID == "" || req.RefreshToken == "" || req.ExpiresAt.IsZero() {
		return nil, errors.New("user_id, refresh_token, and expires_at are required")
	}
	return svc.SessionRepoInterface.Create(ctx, req)
}

func (svc *SessionService) GetByID(ctx context.Context, id string) (*schema.Session, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return svc.SessionRepoInterface.GetByID(ctx, id)
}

func (svc *SessionService) GetByRefreshToken(ctx context.Context, refreshToken string) (*schema.Session, error) {
	if refreshToken == "" {
		return nil, errors.New("refresh_token is required")
	}
	return svc.SessionRepoInterface.GetByRefreshToken(ctx, refreshToken)
}

func (svc *SessionService) Delete(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("id is required")
	}
	return svc.SessionRepoInterface.Delete(ctx, id)
}

func (svc *SessionService) DeleteByUserID(ctx context.Context, userID string) error {
	if userID == "" {
		return errors.New("user_id is required")
	}
	return svc.SessionRepoInterface.DeleteByUserID(ctx, userID)
}
