package model

import (
	"time"

	"github.com/google/uuid"
	"nota.shared/model"
)

type Session struct {
	model.Model
	RefreshToken string    `gorm:"not null;index"`
	ExpiresAt    time.Time `gorm:"not null"`
	userID       uuid.UUID `gorm:"type:uuid;not null;index"`
}
