package main

import (
	"context"
	"log"

	"github.com/djurica-surla/backend-homework/internal/database"
	"github.com/djurica-surla/backend-homework/internal/service"
	"github.com/djurica-surla/backend-homework/internal/storage"
)

func main() {
	connection, err := database.Connect(
		context.Background(),
		database.Config{DSN: "homework.sqlite"},
	)
	if err != nil {
		log.Fatal(err)
	}

	err = database.Migrate(connection, "migrations")
	if err != nil {
		log.Fatal(err)
	}

	questionStorage := storage.NewQuestionStore(connection)

	_ = service.NewQuestionService(questionStorage)
}
