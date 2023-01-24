package storage

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/djurica-surla/backend-homework/internal/entity"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "modernc.org/sqlite"
)

// Represents sqlite implementation of question storage.
type QuestionStore struct {
	db *sql.DB
}

// NewQuestionStore creates a new instance of the QuestionStore.
func NewQuestionStore(connection *sql.DB) *QuestionStore {
	return &QuestionStore{db: connection}
}

// Retrieves a list of questions from the database.
func (store *QuestionStore) GetQuestions(ctx context.Context) ([]entity.Question, error) {
	questions := []entity.Question{}

	rows, err := store.db.QueryContext(ctx,
		`SELECT * FROM question`)
	if err != nil && err != sql.ErrNoRows {
		return []entity.Question{}, fmt.Errorf("error getting questions from db %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		question := entity.Question{}

		err := rows.Scan(
			&question.ID,
			&question.Body,
		)
		if err != nil {
			return nil, fmt.Errorf("error getting questions from database %w", err)
		}

		questions = append(questions, question)
	}

	return questions, nil
}

// Creates a new question in the database.
func (store *QuestionStore) CreateQuestion(ctx context.Context, question entity.Question) error {
	_, err := store.db.ExecContext(ctx,
		`INSERT INTO question (body, correct, question_id)
		VALUES ($1, $2, $3)`, question.Body)
	if err != nil {
		return fmt.Errorf("error creating questions in database %w", err)
	}

	return nil
}

// Updates a question in the database by the id.
func (store *QuestionStore) UpdateQuestion(ctx context.Context, questionID int, body string) error {
	_, err := store.db.ExecContext(ctx,
		`UPDATE question
		SET body = $1 WHERE id = $2`, body, questionID)
	if err != nil {
		return fmt.Errorf("failed to update question")
	}

	return nil
}

// Deletes a question in the database by the id.
func (store *QuestionStore) DeleteQuestion(ctx context.Context, questionID int) error {
	_, err := store.db.ExecContext(ctx,
		`DELETE FROM question
		WHERE id = $1`, questionID)
	if err != nil {
		return fmt.Errorf("failed to delete question")
	}

	return nil
}
