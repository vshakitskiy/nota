package main

import (
	"log"

	"nota.shared/database"
)

func main() {
	_, err := database.ConnectDatabase()
	if err != nil {
		log.Println(err.Error())
		return
	}
}
