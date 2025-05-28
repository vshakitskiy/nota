package repository

import (
	"gorm.io/gorm"
	"nota.shared/config"
)

type Repository struct {
	User    UserRepository
	Session SessionRepository
}

func NewRepository(db *gorm.DB, sessionCfg *config.Session) *Repository {
	return &Repository{
		User:    NewUserRepository(db),
		Session: NewSessionRepository(db, sessionCfg),
	}
}
