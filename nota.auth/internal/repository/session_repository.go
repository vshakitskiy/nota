package repository

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"nota.auth/internal/model"
)

var (
	ErrSessionNotFound = errors.New("session not found")
)

type SessionRepository interface {
	Create(
		ctx context.Context,
		session *model.Session,
	) error
	GetByRefreshToken(ctx context.Context, id uuid.UUID) (*model.Session, error)
	DeleteExpiredByUserID(ctx context.Context, id uuid.UUID) error
}

type SessionRepositoryImpl struct {
	db *gorm.DB
}

func NewSessionRepository(db *gorm.DB) *SessionRepositoryImpl {
	return &SessionRepositoryImpl{
		db: db,
	}
}

func (r *SessionRepositoryImpl) Create(
	ctx context.Context,
	session *model.Session,
) error {
	session.ID = uuid.New()

	if err := r.db.Create(session).Error; err != nil {
		return err
	}

	return nil
}

func (r *SessionRepositoryImpl) GetByRefreshToken(
	ctx context.Context,
	id uuid.UUID,
) (*model.Session, error) {
	panic("not implemented")
}

func (r *SessionRepositoryImpl) DeleteExpiredByUserID(
	ctx context.Context,
	id uuid.UUID,
) error {
	err := r.db.
		Where("user_id = ? and expires_at < ?", id, time.Now()).
		Delete(&model.Session{}).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrSessionNotFound
		}

		return err
	}

	return nil
}
