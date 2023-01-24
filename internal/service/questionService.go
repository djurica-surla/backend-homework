package service

import (
	"context"

	"github.com/djurica-surla/backend-homework/internal/entity"
)

// QuestionStorer represents necessary question storage implementation for question service.
type QuestionStorer interface {
	GetQuestions(ctx context.Context) ([]entity.Question, error)
	CreateQuestion(ctx context.Context, body string) (int, error)
	UpdateQuestion(ctx context.Context, questionID int, body string) error
	DeleteQuestion(ctx context.Context, questionID int) error
}

// QuestionOptionStorer represents necessary question option storage implementation for question service.
type QuestionOptionStorer interface {
	GetQuestionOptions(ctx context.Context, questionID int) ([]entity.QuestionOption, error)
	CreateQuestionOption(ctx context.Context, questionID, correct int, body string) error
	UpdateQuestionOption(ctx context.Context, body string, correct int, questionID int) error
	DeleteQuestionOptions(ctx context.Context, questionID int) error
}

// QuestionService contains business logic for working with question object.
type QuestionService struct {
	questionStore       QuestionStorer
	questionOptionStore QuestionOptionStorer
}

// Instantiates a new question service struct with question repo.
func NewQuestionService(questionStore QuestionStorer, QuestionOptionStore QuestionOptionStorer) *QuestionService {
	return &QuestionService{
		questionStore:       questionStore,
		questionOptionStore: QuestionOptionStore,
	}
}

// GetQuestions handles the logic for getting questions and options from database.
func (s *QuestionService) GetQuestions(ctx context.Context) ([]QuestionDTO, error) {
	questionsEntity, err := s.questionStore.GetQuestions(ctx)
	if err != nil {
		return []QuestionDTO{}, err
	}

	questions := []QuestionDTO{}

	for _, question := range questionsEntity {
		questionOptionsEntity, err := s.questionOptionStore.GetQuestionOptions(ctx, question.ID)
		if err != nil {
			return nil, err
		}

		questionOptions := []QuestionOptionDTO{}

		for _, questionOption := range questionOptionsEntity {
			isCorrect := false

			if questionOption.Correct == 1 {
				isCorrect = true
			}

			questionOptions = append(questionOptions, QuestionOptionDTO{
				questionOption.ID,
				questionOption.Body,
				isCorrect,
			})
		}

		questions = append(questions, QuestionDTO{
			ID:      question.ID,
			Body:    question.Body,
			Options: questionOptions,
		})
	}

	return questions, err
}

// CreateQuestion handles the logic for creating question and its options in database.
func (s *QuestionService) CreateQuestion(ctx context.Context, questionCreation QuestionCreationDTO) error {
	questionID, err := s.questionStore.CreateQuestion(ctx, questionCreation.Body)
	if err != nil {
		return err
	}

	var correctInt int = 0

	for _, option := range questionCreation.Options {
		if option.Correct {
			correctInt = 1
		}

		err := s.questionOptionStore.CreateQuestionOption(
			ctx, questionID, correctInt, option.Body)
		if err != nil {
			return err
		}
	}

	return nil
}
