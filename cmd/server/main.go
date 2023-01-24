package main

import (
	"context"
	"log"

	"github.com/djurica-surla/backend-homework/internal/database"
)

func main() {
	db, err := database.NewRepository(
		context.Background(),
		database.Config{
			DSN: "homework.sqlite",
		})
	if err != nil {
		log.Fatal(err)
	}

	err = db.Migrate("migrations")
	if err != nil {
		log.Fatal(err)
	}
}
