package service

import (
	"context"

	"github.com/djurica-surla/backend-homework/internal/entity"
)

// QuestionStorer represents necessary storage implementation for question service.
type QuestionStorer interface {
	GetQuestions(ctx context.Context) ([]entity.Question, error)
	CreateQuestion(ctx context.Context, question entity.Question) error
	UpdateQuestion(ctx context.Context, questionID int, body string) error
	DeleteQuestion(ctx context.Context, questionID int) error
}

// QuestionService contains business logic for working with question object
type QuestionService struct {
	questionStore QuestionStorer
}

// Instantiates a new question service struct with question repo
func NewQuestionService(questionStore QuestionStorer) *QuestionService {
	return &QuestionService{
		questionStore: questionStore,
	}
}

func (s *QuestionService) GetQuestions(ctx context.Context) ([]Question, error) {
	_, err := s.questionStore.GetQuestions(ctx)
	if err != nil {
		return []Question{}, err
	}
	return []Question{}, err
}
