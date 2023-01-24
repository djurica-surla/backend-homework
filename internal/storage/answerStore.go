package storage

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/djurica-surla/backend-homework/internal/entity"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "modernc.org/sqlite"
)

// Represents sqlite implementation of answer storage.
type AnswerStore struct {
	db *sql.DB
}

// NewAnswerStore creates a new instance of the AnswerStore.
func NewAnswerStore(connection *sql.DB) *AnswerStore {
	return &AnswerStore{db: connection}
}

// Retrieves a list of answers from the database.
func (store *AnswerStore) GetAnswersForQuestion(ctx context.Context, questionID int) ([]entity.Answer, error) {
	answers := []entity.Answer{}

	rows, err := store.db.QueryContext(ctx,
		`SELECT * FROM answer
		WHERE question_id = $1`, questionID)
	if err != nil && err != sql.ErrNoRows {
		return []entity.Answer{}, fmt.Errorf("error getting answers from db %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		answer := entity.Answer{}

		err := rows.Scan(
			&answer.ID,
			&answer.Body,
			&answer.Correct,
			&answer.QuestionID,
		)
		if err != nil {
			return nil, fmt.Errorf("error getting answers from database %w", err)
		}

		answers = append(answers, answer)
	}

	return answers, nil
}

// Creates a new Answer in the database.
func (store *AnswerStore) CreateAnswer(ctx context.Context, answer entity.Answer) error {
	_, err := store.db.ExecContext(ctx,
		`INSERT INTO answer (body, correct, question_id)
		VALUES ($1, $2, $3)`, answer.Body, answer.Correct, answer.QuestionID)
	if err != nil {
		return fmt.Errorf("error creating answers in database %w", err)
	}

	return nil
}

// Updates an Answer in the database by the question id.
func (store *AnswerStore) UpdateAnswerForQuestion(ctx context.Context,
	body string, correct int, questionID int) error {
	_, err := store.db.ExecContext(ctx,
		`UPDATE answer
		SET body = $1, correct = $2 WHERE question_id = $2`, body, correct, questionID)
	if err != nil {
		return fmt.Errorf("failed to update answer")
	}

	return nil
}

// Deletes a Answer in the database by the question id.
func (store *AnswerStore) DeleteAnswersForQuestion(ctx context.Context, questionID int) error {
	_, err := store.db.ExecContext(ctx,
		`DELETE FROM Answer
		WHERE question_id = $1`, questionID)
	if err != nil {
		return fmt.Errorf("failed to delete answer")
	}

	return nil
}
