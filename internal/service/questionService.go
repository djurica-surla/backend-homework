package service

import (
	"context"
	"fmt"

	"github.com/djurica-surla/backend-homework/internal/entity"
)

// QuestionStorer represents necessary question storage implementation for question service.
type QuestionStorer interface {
	GetQuestions(ctx context.Context, pageSize, offset int) ([]entity.Question, error)
	GetQuestionByID(ctx context.Context, questionID int) (entity.Question, error)
	CreateQuestion(ctx context.Context, body string) (int, error)
	UpdateQuestion(ctx context.Context, questionID int, body string) (int, error)
	DeleteQuestion(ctx context.Context, questionID int) error
}

// QuestionOptionStorer represents necessary question option storage implementation for question service.
type QuestionOptionStorer interface {
	GetQuestionOptions(ctx context.Context, questionID int) ([]entity.QuestionOption, error)
	CreateQuestionOption(ctx context.Context, questionID int, questionOption QuestionOptionCreationDTO) error
	// UpdateQuestionOption(ctx context.Context, body string, correct int, questionID int) error
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

// GetQuestions handles the logic for getting questions and its options.
func (s *QuestionService) GetQuestions(ctx context.Context, pageSize, offset int) ([]QuestionDTO, error) {
	questionsEntity, err := s.questionStore.GetQuestions(ctx, pageSize, offset)
	if err != nil {
		return nil, err
	}

	questions := []QuestionDTO{}

	for _, question := range questionsEntity {
		questionOptionsEntity, err := s.questionOptionStore.GetQuestionOptions(ctx, question.ID)
		if err != nil {
			return nil, err
		}

		questionOptions := []QuestionOptionDTO{}

		for _, questionOption := range questionOptionsEntity {
			questionOptions = append(questionOptions, QuestionOptionDTO{
				questionOption.ID,
				questionOption.Body,
				questionOption.Correct,
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

// GetQuestions handles the logic for getting questions by id.
func (s *QuestionService) GetQuestionByID(ctx context.Context, questionID int) (QuestionDTO, error) {
	questionEntity, err := s.questionStore.GetQuestionByID(ctx, questionID)
	if err != nil {
		return QuestionDTO{}, err
	}

	questionOptionsEntity, err := s.questionOptionStore.GetQuestionOptions(ctx, questionEntity.ID)
	if err != nil {
		return QuestionDTO{}, err
	}

	questionOptions := []QuestionOptionDTO{}

	for _, questionOption := range questionOptionsEntity {
		questionOptions = append(questionOptions, QuestionOptionDTO{
			questionOption.ID,
			questionOption.Body,
			questionOption.Correct,
		})
	}

	question := QuestionDTO{
		ID:      questionEntity.ID,
		Body:    questionEntity.Body,
		Options: questionOptions,
	}

	return question, err
}

// CreateQuestion handles the logic for creating question and its options in database.
func (s *QuestionService) CreateQuestion(ctx context.Context, questionCreation QuestionCreationDTO) (QuestionDTO, error) {
	questionID, err := s.questionStore.CreateQuestion(ctx, questionCreation.Body)
	if err != nil {
		return QuestionDTO{}, err
	}

	for _, option := range questionCreation.Options {
		err := s.questionOptionStore.CreateQuestionOption(ctx, questionID, option)
		if err != nil {
			return QuestionDTO{}, err
		}
	}

	// Retrieve the new records.
	questionDTO, err := s.GetQuestionByID(ctx, questionID)
	if err != nil {
		return QuestionDTO{}, fmt.Errorf("error trying to update question: %w", err)
	}

	return questionDTO, nil
}

// UpdateQuestion handles the logic for updating question and its options in database.
func (s *QuestionService) UpdateQuestion(ctx context.Context,
	questionID int, questionCreation QuestionCreationDTO) (QuestionDTO, error) {
	// Update the question record first
	rowsAffected, err := s.questionStore.UpdateQuestion(ctx, questionID, questionCreation.Body)
	if err != nil {
		return QuestionDTO{}, err
	}

	// If rows affected are zero, return empty dto meaning nothing was updated
	if rowsAffected == 0 {
		return QuestionDTO{}, err
	}

	// Delete the previous options since we are replacing them.
	err = s.questionOptionStore.DeleteQuestionOptions(ctx, questionID)
	if err != nil {
		return QuestionDTO{}, fmt.Errorf("error trying to update question: %w", err)
	}

	for _, option := range questionCreation.Options {
		// Insert new options into the database.
		err = s.questionOptionStore.CreateQuestionOption(ctx, questionID, option)
		if err != nil {
			return QuestionDTO{}, fmt.Errorf("error trying to update question: %w", err)
		}
	}

	// Retrieve the new records.
	questionDTO, err := s.GetQuestionByID(ctx, questionID)
	if err != nil {
		return QuestionDTO{}, fmt.Errorf("error trying to update question: %w", err)
	}

	return questionDTO, nil
}

// DeleteQuestion handles the logic for deleting question and its options in database.
func (s *QuestionService) DeleteQuestion(ctx context.Context, questionID int) error {
	return s.questionStore.DeleteQuestion(ctx, questionID)
}
