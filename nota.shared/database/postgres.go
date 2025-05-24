package database

import (
	"errors"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/plugin/opentelemetry/tracing"
	"nota.shared/config"
)

func ConnectDatabase() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Moscow",
		config.GetenvDefault("POSTGRES_HOST", "localhost"),
		config.GetenvDefault("POSTGRES_USER", "admeanie"),
		config.GetenvDefault("POSTGRES_PASSWORD", "shabi"),
		config.GetenvDefault("POSTGRES_DB", "nota"),
		config.GetenvDefault("POSTGRES_PORT", "4040"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, errors.New("unable to connect to database")
	}

	if err := db.Use(tracing.NewPlugin()); err != nil {
		return nil, errors.New("unable to use database tracing plugin")
	}

	db.Migrator()

	return db, nil
}

func Migrate(db *gorm.DB, dst ...interface{}) error {
	if err := db.AutoMigrate(dst...); err != nil {
		return errors.New("unable to migrate")
	}

	return nil
}
