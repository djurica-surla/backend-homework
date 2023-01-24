package main

import (
	"context"
	"log"

	"github.com/djurica-surla/backend-homework/internal/database"
	"github.com/djurica-surla/backend-homework/internal/repository/sqlite"
)

func main() {
	connection, err := database.Connect(
		context.Background(),
		database.Config{DSN: "homework.sqlite"},
	)
	if err != nil {
		log.Fatal(err)
	}

	err = database.Migrate(connection, "/migrations")
	if err != nil {
		log.Fatal(err)
	}

	_ = sqlite.NewQuestionRepository(connection)
}
