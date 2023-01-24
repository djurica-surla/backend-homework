package sqlite

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/djurica-surla/backend-homework/internal/entity"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "modernc.org/sqlite"
)

// Sqlite implementation of QuestionRepo interface.
type QuestionRepo struct {
	db *sql.DB
}

// NewQuestionRepository creates a new instance of the QuestionRepo.
func NewQuestionRepository(connection *sql.DB) QuestionRepo {
	return QuestionRepo{db: connection}
}

// Retrieves a list of questions from the database.
func (repo QuestionRepo) GetQuestions(ctx context.Context) ([]entity.Question, error) {
	questions := []entity.Question{}

	rows, err := repo.db.QueryContext(ctx,
		`SELECT * FROM question`)
	if err != nil && err != sql.ErrNoRows {
		return []entity.Question{}, fmt.Errorf("failed to get questions")
	}
	defer rows.Close()

	for rows.Next() {
		question := entity.Question{}

		err := rows.Scan(
			&question.ID,
			&question.Body,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to get questions")
		}

		questions = append(questions, question)
	}

	return questions, nil
}

// Creates a new question in the database.
func (repo QuestionRepo) CreateQuestion(ctx context.Context, question entity.Question) error {
	err := repo.db.QueryRowContext(ctx,
		`INSERT INTO answer (body, correct, question_id)
		VALUES ($1, $2, $3)`, question.Body)
	if err != nil {
		return fmt.Errorf("failed to create question")
	}

	return nil
}

// Updates a question in the database by the id.
func (repo QuestionRepo) UpdateQuestion(ctx context.Context, questionID int, body string) error {
	err := repo.db.QueryRowContext(ctx,
		`UPDATE question
		SET body = $1 WHERE id = $2`, body, questionID)
	if err != nil {
		return fmt.Errorf("failed to update question")
	}

	return nil
}

// Deletes a question in the database by the id.
func (repo QuestionRepo) DeleteQuestion(ctx context.Context, questionID int) error {
	err := repo.db.QueryRowContext(ctx,
		`DELETE FROM question
		WHERE id = $1`, questionID)
	if err != nil {
		return fmt.Errorf("failed to delete question")
	}

	return nil
}
