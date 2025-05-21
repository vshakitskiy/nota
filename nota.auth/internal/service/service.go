package service

import "nota.auth/internal/repository"

type Service struct {
	Auth   AuthService
	Access AccessService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Auth:   NewAuthService(repo),
		Access: NewAccessService(repo),
	}
}
