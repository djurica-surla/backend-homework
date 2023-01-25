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
func (store *QuestionStore) GetQuestions(ctx context.Context, pageSize, offset int) ([]entity.Question, error) {
	questions := []entity.Question{}

	rows, err := store.db.QueryContext(ctx,
		`SELECT * FROM question 
		  LIMIT $2 OFFSET $1`, offset, pageSize)
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

// Retrieves a  question from database the id.
func (store *QuestionStore) GetQuestionByID(ctx context.Context, questionID int) (entity.Question, error) {
	question := entity.Question{}

	err := store.db.QueryRowContext(ctx,
		`SELECT * FROM question WHERE id = $1`, questionID).
		Scan(&question.ID, &question.Body)
	if err != nil {
		return entity.Question{}, fmt.Errorf("error getting question from db %w", err)
	}

	return question, nil
}

// Creates a new question in the database.
func (store *QuestionStore) CreateQuestion(ctx context.Context, body string) (int, error) {
	var questionID int

	err := store.db.QueryRowContext(ctx,
		`INSERT INTO question (body)
		VALUES ($1) RETURNING id`, body).Scan(&questionID)
	if err != nil {
		return 0, fmt.Errorf("error creating questions in database %w", err)
	}

	return questionID, nil
}

// Updates a question in the database by the id.
func (store *QuestionStore) UpdateQuestion(ctx context.Context, questionID int, body string) error {
	_, err := store.db.ExecContext(ctx,
		`UPDATE question
		SET body = $1 WHERE id = $2`, body, questionID)
	if err != nil {
		return fmt.Errorf("failed to update question %w", err)
	}

	return nil
}

// Deletes a question in the database by the id.
func (store *QuestionStore) DeleteQuestion(ctx context.Context, questionID int) error {
	_, err := store.db.ExecContext(ctx,
		`DELETE FROM question
		WHERE id = $1`, questionID)
	if err != nil {
		return fmt.Errorf("failed to delete question %w", err)
	}

	return nil
}
