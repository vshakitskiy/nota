package repository

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"nota.auth/internal/model"
)

type SessionRepository interface {
	Create(
		ctx context.Context,
		session *model.Session,
	) error
	GetByRefreshToken(ctx context.Context, id uuid.UUID) (*model.Session, error)
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
	id uuid.UUID,
) (*model.Session, error) {
	panic("not implemented")
}

func (r *SessionRepositoryImpl) DeleteByUserID(
	ctx context.Context,
	id uuid.UUID,
) error {
	panic("not implemented")
}
