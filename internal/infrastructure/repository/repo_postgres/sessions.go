package repo_postgres

import (
	"context"
	"github.com/Rafiur/wallet_app/internal/infrastructure/repository/schema"
	"github.com/uptrace/bun"
)

type SessionRepo struct {
	db *bun.DB
}

func NewSessionRepo(db *bun.DB) *SessionRepo {
	return &SessionRepo{db: db}
}

func (repo *SessionRepo) Create(ctx context.Context, req *schema.Session) (*schema.Session, error) {
	_, err := repo.db.NewInsert().
		Model(req).
		Returning("*").
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (repo *SessionRepo) GetByID(ctx context.Context, id string) (*schema.Session, error) {
	var session schema.Session
	err := repo.db.NewSelect().
		Model(&session).
		Where("id = ?", id).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return &session, nil
}

func (repo *SessionRepo) GetByRefreshToken(ctx context.Context, refreshToken string) (*schema.Session, error) {
	var session schema.Session
	err := repo.db.NewSelect().
		Model(&session).
		Where("refresh_token = ?", refreshToken).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return &session, nil
}

func (repo *SessionRepo) Delete(ctx context.Context, id string) error {
	_, err := repo.db.NewDelete().
		Model((*schema.Session)(nil)).
		Where("id = ?", id).
		Exec(ctx)
	return err
}

func (repo *SessionRepo) DeleteByUserID(ctx context.Context, userID string) error {
	_, err := repo.db.NewDelete().
		Model((*schema.Session)(nil)).
		Where("user_id = ?", userID).
		Exec(ctx)
	return err
}

// Optionally update expiration date or token
//func (repo *SessionRepo) Update(ctx context.Context, req *schema.Session) (*schema.Session, error) {
//	existing, err := repo.GetByID(ctx, req.ID)
//	if err != nil {
//		return nil, err
//	}
//
//	if req.RefreshToken != "" && req.RefreshToken != existing.RefreshToken {
//		existing.RefreshToken = req.RefreshToken
//	}
//	if !req.ExpiresAt.IsZero() {
//		existing.ExpiresAt = req.ExpiresAt
//	}
//
//	_, err = repo.db.NewUpdate().
//		Model(existing).
//		Where("id = ?", req.ID).
//		Exec(ctx)
//	if err != nil {
//		return nil, err
//	}
//	return existing, nil
//}
