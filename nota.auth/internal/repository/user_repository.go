package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"nota.auth/internal/model"
)

var (
	ErrUserAlreadyExists = errors.New("user already exists")
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) (*uuid.UUID, error)
	GetById(ctx context.Context, id uuid.UUID) (*model.User, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (r *UserRepositoryImpl) Create(
	ctx context.Context,
	user *model.User,
) (*uuid.UUID, error) {
	var exists bool
	err := r.db.
		Table("users").
		Select("count(*) > 0").
		Where("email = ?", user.Email).
		Find(&exists).Error
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrUserAlreadyExists
	}

	if err := r.db.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user.ID, nil
}

func (r *UserRepositoryImpl) GetById(
	ctx context.Context,
	id uuid.UUID,
) (*model.User, error) {
	return nil, nil
}

func (r *UserRepositoryImpl) GetByEmail(
	ctx context.Context,
	email string,
) (*model.User, error) {
	return nil, nil
}
