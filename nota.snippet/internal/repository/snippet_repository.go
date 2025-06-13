package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"nota.shared/pagination"
	"nota.snippet/internal/model"
)

var ErrSnippetNotFound = errors.New("snippet not found")

type SnippetRepository interface {
	Create(ctx context.Context, snippet *model.Snippet) (*uuid.UUID, error)
	GetByID(
		ctx context.Context,
		id uuid.UUID,
	) (*model.Snippet, error)
	List(ctx context.Context, pagination pagination.Pagination) ([]*model.Snippet, error)
	ListByOwnerID(
		ctx context.Context,
		ownerID uuid.UUID,
		pagination pagination.Pagination,
	) ([]*model.Snippet, int64, error)
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
	snippet.ID = uuid.New()
	snippet.CreatedAt = time.Now()
	snippet.UpdatedAt = time.Now()

	fmt.Println("snippet", snippet)

	if err := r.db.Create(snippet).Error; err != nil {
		return nil, err
	}

	return &snippet.ID, nil
}

func (r *snippetRepository) GetByID(
	ctx context.Context,
	id uuid.UUID,
) (*model.Snippet, error) {
	var snippet model.Snippet
	if err := r.db.WithContext(ctx).First(&snippet, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrSnippetNotFound
		}

		return nil, err
	}

	return &snippet, nil
}

func (r *snippetRepository) List(
	ctx context.Context,
	pagination pagination.Pagination,
) ([]*model.Snippet, error) {
	panic("not implemented")
}

func (r *snippetRepository) ListByOwnerID(
	ctx context.Context,
	ownerID uuid.UUID,
	pagination pagination.Pagination,
) ([]*model.Snippet, int64, error) {
	var total int64

	err := r.db.
		WithContext(ctx).
		Model(&model.Snippet{}).
		Where("owner_id = ?", ownerID).
		Count(&total).Error
	if err != nil {
		return nil, total, err
	}

	var snippets []*model.Snippet
	err = r.db.
		WithContext(ctx).
		Where("owner_id = ?", ownerID).
		Order("created_at DESC").
		Offset((pagination.Page - 1) * pagination.Size).
		Limit(pagination.Size).
		Find(&snippets).Error
	if err != nil {
		return nil, total, err
	}

	return snippets, total, nil
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
