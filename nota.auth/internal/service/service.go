package service

import "nota.auth/internal/repository"

type Service struct {
	Auth AuthService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Auth: NewAuthService(repo),
	}
}
