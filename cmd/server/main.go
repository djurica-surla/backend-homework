package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/djurica-surla/backend-homework/internal/database"
	"github.com/djurica-surla/backend-homework/internal/service"
	"github.com/djurica-surla/backend-homework/internal/storage"
	transporthttp "github.com/djurica-surla/backend-homework/internal/transport/http"
	"github.com/gorilla/mux"
)

func main() {
	// Attempt to establish a connection with the database.
	connection, err := database.Connect(
		context.Background(),
		database.Config{DSN: "homework.sqlite"},
	)
	if err != nil {
		log.Fatal(err)
	}

	// Run up migrations to create database schema.
	err = database.Migrate(connection, "migrations")
	if err != nil {
		log.Fatal(err)
	}

	// Instantiate question storage.
	questionStorage := storage.NewQuestionStore(connection)

	// Instantiate question option storage.
	questionOptionStorage := storage.NewQuestionOptionStore(connection)

	// Instantiate question service.
	questionService := service.NewQuestionService(questionStorage, questionOptionStorage)

	// Instantiate mux router.
	router := mux.NewRouter().StrictSlash(true)

	// Instantiate question handler.
	handler := transporthttp.NewQuestionHandler(questionService)

	// Register routes for question handler.
	handler.RegisterRoutes(router)

	// Start the server
	log.Printf("Starting Server on port %s", "8080")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", "8080"), router))
}
