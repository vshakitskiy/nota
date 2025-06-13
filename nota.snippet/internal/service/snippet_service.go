package service

import (
	"context"

	"github.com/google/uuid"
	sharedModel "nota.shared/model"
	"nota.snippet/internal/model"
	"nota.snippet/internal/repository"
)

type SnippetService interface {
	Create(ctx context.Context, snippet *model.Snippet) (*uuid.UUID, error)
	List(ctx context.Context, pagination sharedModel.Pagination) ([]*model.Snippet, error)
	GetByID(
		ctx context.Context,
		id uuid.UUID,
		pagination sharedModel.Pagination,
	) (*model.Snippet, error)
	GetByOwnerID(
		ctx context.Context,
		ownerID uuid.UUID, pagination sharedModel.Pagination) ([]*model.Snippet, error)
	Update(ctx context.Context, id uuid.UUID, snippet *model.Snippet) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type snippetService struct {
	repo *repository.Repository
}

func NewSnippetService(repo *repository.Repository) *snippetService {
	return &snippetService{repo: repo}
}

func (s *snippetService) Create(
	ctx context.Context,
	snippet *model.Snippet,
) (*uuid.UUID, error) {
	panic("not implemented")
}

func (s *snippetService) List(
	ctx context.Context,
	pagination sharedModel.Pagination,
) ([]*model.Snippet, error) {
	panic("not implemented")
}

func (s *snippetService) GetByID(
	ctx context.Context,
	id uuid.UUID,
	pagination sharedModel.Pagination,
) (*model.Snippet, error) {
	panic("not implemented")
}

func (s *snippetService) GetByOwnerID(
	ctx context.Context,
	ownerID uuid.UUID,
	pagination sharedModel.Pagination,
) ([]*model.Snippet, error) {
	panic("not implemented")
}

func (s *snippetService) Update(
	ctx context.Context,
	id uuid.UUID,
	snippet *model.Snippet,
) error {
	panic("not implemented")
}

func (s *snippetService) Delete(ctx context.Context, id uuid.UUID) error {
	panic("not implemented")
}
