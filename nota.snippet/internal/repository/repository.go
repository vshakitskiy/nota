package repository

import "gorm.io/gorm"

type Repository struct {
	Snippet SnippetRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Snippet: NewSnippetRepository(db),
	}
}
