package repository

import "gorm.io/gorm"

type Repository struct {
	User    UserRepository
	Session SessionRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		User:    NewUserRepository(db),
		Session: NewSessionRepository(db),
	}
}
