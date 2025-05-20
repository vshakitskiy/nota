package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"nota.auth/internal/model"
	"nota.auth/internal/repository"
	"nota.auth/pkg/bcrypt"
)

var (
	ErrInvalidPassword = errors.New("password is invalid")
)

type UserService interface {
	Create(ctx context.Context, username, email, password string) (*uuid.UUID, error)
}

type UserServiceImpl struct {
	repo *repository.Repository
}

func NewUserService(repo *repository.Repository) *UserServiceImpl {
	return &UserServiceImpl{
		repo: repo,
	}
}

func (s *UserServiceImpl) Create(
	ctx context.Context,
	username, email, password string,
) (*uuid.UUID, error) {
	user := new(model.User)
	user.ID = uuid.New()
	user.Username = username
	user.Email = email

	passwordHash, err := bcrypt.HashPassword(password)
	if err != nil {
		return nil, ErrInvalidPassword
	}
	user.Password = passwordHash

	return s.repo.User.Create(ctx, user)
}
