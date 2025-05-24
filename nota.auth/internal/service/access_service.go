package service

import (
	"context"

	"github.com/google/uuid"
	"nota.auth/internal/repository"
	"nota.auth/pkg/jwt"
	"nota.shared/telemetry"
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
	ctx, span := telemetry.StartSpan(ctx, "AccessService.Validate")
	defer span.End()

	claims, err := jwt.ValidateJWT(accessToken)
	if err != nil {
		telemetry.RecordError(span, err)
		return nil, err
	}

	return &claims.UserID, nil
}
