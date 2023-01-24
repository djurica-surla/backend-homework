package repository

import (
	"context"

	"github.com/djurica-surla/backend-homework/internal/entity"
)

// QuestionRepository is an interface which defines methods that question repo needs to implement.
type QuestionRepository interface {
	GetQuestions(ctx context.Context) ([]entity.Question, error)
	CreateQuestion(ctx context.Context, question entity.Question) error
	UpdateQuestion(ctx context.Context, questionID int, body string) error
	DeleteQuestion(ctx context.Context, questionID int) error
}
