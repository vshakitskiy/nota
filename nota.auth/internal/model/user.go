package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	ID       uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Username string    `gorm:"UniqueIndex;not null"`
	Email    string    `gorm:"UniqueIndex;not null"`
	Password string
}
