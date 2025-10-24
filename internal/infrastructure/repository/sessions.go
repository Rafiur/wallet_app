package repository

import (
	"context"
	"github.com/Rafiur/wallet_app/internal/infrastructure/repository/schema"
)

type SessionRepoInterface interface {
	Create(ctx context.Context, req *schema.Session) (*schema.Session, error)
	GetByID(ctx context.Context, id string) (*schema.Session, error)
	GetByRefreshToken(ctx context.Context, refreshToken string) (*schema.Session, error)
	Delete(ctx context.Context, id string) error
	DeleteByUserID(ctx context.Context, userID string) error
}
