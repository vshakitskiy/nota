package model

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
	"nota.shared/model"
)

type Snippet struct {
	model.Model

	Title        string         `gorm:"not null"`
	Content      string         `gorm:"not null"`
	LanguageHint string         `gorm:"not null;default:'text'"`
	IsPublic     bool           `gorm:"not null;default:false"`
	Tags         pq.StringArray `gorm:"type:text[]"`
	OwnerID      uuid.UUID      `gorm:"not null"`
}
