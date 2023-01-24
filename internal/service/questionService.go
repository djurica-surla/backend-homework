package service

import (
	"context"

	"github.com/djurica-surla/backend-homework/internal/entity"
)

// QuestionStorer represents necessary question storage implementation for question service.
type QuestionStorer interface {
	GetQuestions(ctx context.Context) ([]entity.Question, error)
	CreateQuestion(ctx context.Context, question entity.Question) error
	UpdateQuestion(ctx context.Context, questionID int, body string) error
	DeleteQuestion(ctx context.Context, questionID int) error
}

// QuestionOptionStorer represents necessary question option storage implementation for question service.
type QuestionOptionStorer interface {
	GetQuestionOptions(ctx context.Context, questionID int) ([]entity.QuestionOption, error)
	CreateQuestionOption(ctx context.Context, QuestionOption entity.QuestionOption) error
	UpdateQuestionOption(ctx context.Context, body string, correct int, questionID int) error
	DeleteQuestionOptions(ctx context.Context, questionID int) error
}

// QuestionService contains business logic for working with question object.
type QuestionService struct {
	questionStore       QuestionStorer
	QuestionOptionStore QuestionOptionStorer
}

// Instantiates a new question service struct with question repo.
func NewQuestionService(questionStore QuestionStorer, QuestionOptionStore QuestionOptionStorer) *QuestionService {
	return &QuestionService{
		questionStore:       questionStore,
		QuestionOptionStore: QuestionOptionStore,
	}
}

// QuestionService handles the logic for getting questions and options from database.
func (s *QuestionService) GetQuestions(ctx context.Context) ([]Question, error) {
	questionsEntity, err := s.questionStore.GetQuestions(ctx)
	if err != nil {
		return []Question{}, err
	}

	questions := []Question{}
	questionOptions := []QuestionOption{}

	for _, question := range questionsEntity {
		questionOptionsEntity, err := s.QuestionOptionStore.GetQuestionOptions(ctx, question.ID)
		if err != nil {
			return nil, err
		}

		for _, questionOption := range questionOptionsEntity {
			questionOptions = append(questionOptions, QuestionOption{
				questionOption.ID,
				questionOption.Body,
				questionOption.Correct,
			})
		}

		questions = append(questions, Question{
			ID:      question.ID,
			Body:    question.Body,
			Options: questionOptions,
		})
	}

	return questions, err
}
