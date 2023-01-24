package storage

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/djurica-surla/backend-homework/internal/entity"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "modernc.org/sqlite"
)

// Represents sqlite implementation of QuestionOption storage.
type QuestionOptionStore struct {
	db *sql.DB
}

// NewQuestionOptionStore creates a new instance of the QuestionOptionStore.
func NewQuestionOptionStore(connection *sql.DB) *QuestionOptionStore {
	return &QuestionOptionStore{db: connection}
}

// Retrieves a list of options for a questions from the database.
func (store *QuestionOptionStore) GetQuestionOptions(ctx context.Context, questionID int) ([]entity.QuestionOption, error) {
	questionOptions := []entity.QuestionOption{}

	rows, err := store.db.QueryContext(ctx,
		`SELECT * FROM question_option
		WHERE question_id = $1`, questionID)
	if err != nil && err != sql.ErrNoRows {
		return nil, fmt.Errorf("error getting question options from db %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		questionOption := entity.QuestionOption{}

		err := rows.Scan(
			&questionOption.ID,
			&questionOption.Body,
			&questionOption.Correct,
			&questionOption.QuestionID,
		)
		if err != nil {
			return nil, fmt.Errorf("error getting question options from database %w", err)
		}

		questionOptions = append(questionOptions, questionOption)
	}

	return questionOptions, nil
}

// Creates a new QuestionOption in the database.
func (store *QuestionOptionStore) CreateQuestionOption(ctx context.Context,
	questionID, correct int, body string) error {
	_, err := store.db.ExecContext(ctx,
		`INSERT INTO question_option (body, correct, question_id)
		VALUES ($1, $2, $3)`, body, correct, questionID)
	if err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("error creating question options in database %w", err)
	}

	return nil
}

// Updates an QuestionOption in the database by the question id.
func (store *QuestionOptionStore) UpdateQuestionOption(ctx context.Context,
	body string, correct int, questionID int) error {
	_, err := store.db.ExecContext(ctx,
		`UPDATE QuestionOption
		SET body = $1, correct = $2 WHERE question_id = $2`, body, correct, questionID)
	if err != nil {
		return fmt.Errorf("failed to update question option")
	}

	return nil
}

// Deletes a QuestionOption in the database by the question id.
func (store *QuestionOptionStore) DeleteQuestionOptions(ctx context.Context, questionID int) error {
	_, err := store.db.ExecContext(ctx,
		`DELETE FROM QuestionOption
		WHERE question_id = $1`, questionID)
	if err != nil {
		return fmt.Errorf("failed to delete question option")
	}

	return nil
}
