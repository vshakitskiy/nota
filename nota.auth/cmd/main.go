package main

import (
	"context"
	"log"

	. "nota.auth/internal/app"
	"nota.auth/internal/model"
	"nota.shared/config"
	"nota.shared/database"
)

func main() {
	if err := config.LoadEnv(".env"); err != nil {
		log.Fatal(err.Error())
	}

	db, err := database.ConnectDatabase()
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Running migration..")
	if err = database.Migrate(db, model.User{}, model.Session{}); err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Starting app...")
	app := NewApp(db)
	ctx := context.Background()

	if err := app.Start(ctx); err != nil {
		log.Fatal(err.Error())
	}
}
