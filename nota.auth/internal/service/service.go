package service

import (
	"nota.auth/internal/repository"
	"nota.shared/config"
)

type Service struct {
	Auth   AuthService
	Access AccessService
}

func NewService(repo *repository.Repository, cfg *config.Jwt) *Service {
	return &Service{
		Auth:   NewAuthService(repo, cfg),
		Access: NewAccessService(repo),
	}
}
