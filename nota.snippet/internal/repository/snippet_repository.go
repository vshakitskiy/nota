package repository

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
	sharedModel "nota.shared/model"
	"nota.snippet/internal/model"
)

type SnippetRepository interface {
	Create(ctx context.Context, snippet *model.Snippet) (*uuid.UUID, error)
	List(ctx context.Context, pagination sharedModel.Pagination) ([]*model.Snippet, error)
	GetByID(
		ctx context.Context,
		id uuid.UUID,
		pagination sharedModel.Pagination,
	) (*model.Snippet, error)
	GetByOwnerID(
		ctx context.Context,
		ownerID uuid.UUID,
		pagination sharedModel.Pagination,
	) ([]*model.Snippet, error)
	Update(
		ctx context.Context,
		id uuid.UUID,
		snippet *model.Snippet,
	) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type snippetRepository struct {
	db *gorm.DB
}

func NewSnippetRepository(db *gorm.DB) *snippetRepository {
	return &snippetRepository{db: db}
}

func (r *snippetRepository) Create(
	ctx context.Context,
	snippet *model.Snippet,
) (*uuid.UUID, error) {
	panic("not implemented")
}

func (r *snippetRepository) List(
	ctx context.Context,
	pagination sharedModel.Pagination,
) ([]*model.Snippet, error) {
	panic("not implemented")
}

func (r *snippetRepository) GetByID(
	ctx context.Context,
	id uuid.UUID,
	pagination sharedModel.Pagination,
) (*model.Snippet, error) {
	panic("not implemented")
}

func (r *snippetRepository) GetByOwnerID(
	ctx context.Context,
	ownerID uuid.UUID,
	pagination sharedModel.Pagination,
) ([]*model.Snippet, error) {
	panic("not implemented")
}

func (r *snippetRepository) Update(
	ctx context.Context,
	id uuid.UUID,
	snippet *model.Snippet,
) error {
	panic("not implemented")
}

func (r *snippetRepository) Delete(ctx context.Context, id uuid.UUID) error {
	panic("not implemented")
}
