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
	GetByRefreshToken(ctx context.Context, refreshToken string) (*model.Session, error)
	DeleteExpiredByUserID(ctx context.Context, id uuid.UUID) error
	DeleteByUserID(ctx context.Context, id uuid.UUID) error
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
	refreshToken string,
) (*model.Session, error) {
	session := new(model.Session)

	if err := r.db.Where("refresh_token = ?", refreshToken).First(session).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrSessionNotFound
		}

		return nil, err
	}

	return session, nil
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
		return err
	}

	return nil
}

func (r *SessionRepositoryImpl) DeleteByUserID(
	ctx context.Context,
	id uuid.UUID,
) error {
	err := r.db.Where("user_id = ?", id).Delete(&model.Session{}).Error
	if err != nil {
		return err
	}

	return nil
}
