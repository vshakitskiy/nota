package model

import (
	"nota.shared/model"
)

type User struct {
	model.Model
	Username string `gorm:"not null"`
	Email    string `gorm:"UniqueIndex;not null"`
	Password string `gorm:"column:password_hash;not null"`
}
