package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"nota.auth/internal/model"
	"nota.shared/telemetry"
)

var (
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrUserNotFound      = errors.New("user not found")
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
	ctx, span := telemetry.StartSpan(ctx, "UserRepository.Create")
	defer span.End()

	user.ID = uuid.New()

	var exists bool
	err := r.db.
		Model(&model.User{}).
		Select("count(*) > 0").
		Where("email = ?", user.Email).
		Find(&exists).Error
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrUserAlreadyExists
	}

	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}

	return &user.ID, nil
}

func (r *UserRepositoryImpl) GetById(
	ctx context.Context,
	id uuid.UUID,
) (*model.User, error) {
	ctx, span := telemetry.StartSpan(ctx, "UserRepository.GetById")
	defer span.End()

	user := new(model.User)
	err := r.db.
		Where("id = ?", id).
		First(user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		} else {
			return nil, err
		}
	}

	return user, nil
}

func (r *UserRepositoryImpl) GetByEmail(
	ctx context.Context,
	email string,
) (*model.User, error) {
	ctx, span := telemetry.StartSpan(ctx, "UserRepository.GetByEmail")
	defer span.End()

	user := new(model.User)
	err := r.db.
		Where("email = ?", email).
		First(user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		} else {
			return nil, err
		}
	}

	return user, nil
}
