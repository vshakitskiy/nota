package service

import "nota.snippet/internal/repository"

type Service struct {
	Snippet SnippetService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Snippet: NewSnippetService(repo),
	}
}
