package service

import (
	"context"

	"github.com/google/uuid"
	"nota.auth/internal/repository"
	"nota.auth/pkg/jwt"
)

type AccessService interface {
	Validate(ctx context.Context, accessToken string) (*uuid.UUID, error)
}

type AccessServiceImpl struct {
	repo *repository.Repository
}

func NewAccessService(repo *repository.Repository) *AccessServiceImpl {
	return &AccessServiceImpl{
		repo: repo,
	}
}

func (s *AccessServiceImpl) Validate(ctx context.Context, accessToken string) (*uuid.UUID, error) {
	claims, err := jwt.ValidateJWT(accessToken)
	if err != nil {
		return nil, err
	}

	return &claims.UserID, nil
}
