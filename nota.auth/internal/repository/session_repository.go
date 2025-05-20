package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"nota.auth/internal/model"
)

type SessionRepository interface {
	Create(
		ctx context.Context,
		userID uuid.UUID,
		refreshToken string,
		expiresAt time.Time,
	) error
	GetByRefreshToken(ctx context.Context, id uuid.UUID) (*model.Session, error)
	DeleteByUserID(ctx context.Context, id uuid.UUID) error
}
