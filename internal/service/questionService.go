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

// QuestionStorer represents necessary answer storage implementation for question service.
type AnswerStorer interface {
	GetAnswersForQuestion(ctx context.Context, questionID int) ([]entity.Answer, error)
	CreateAnswer(ctx context.Context, answer entity.Answer) error
	UpdateAnswerForQuestion(ctx context.Context, body string, correct int, questionID int) error
	DeleteAnswersForQuestion(ctx context.Context, questionID int) error
}

// QuestionService contains business logic for working with question object
type QuestionService struct {
	questionStore QuestionStorer
	answerStore   AnswerStorer
}

// Instantiates a new question service struct with question repo
func NewQuestionService(questionStore QuestionStorer, answerStore AnswerStorer) *QuestionService {
	return &QuestionService{
		questionStore: questionStore,
		answerStore:   answerStore,
	}
}

// Retrieves questions with nested answers from the database
func (s *QuestionService) GetQuestions(ctx context.Context) ([]Question, error) {
	questionsEntity, err := s.questionStore.GetQuestions(ctx)
	if err != nil {
		return []Question{}, err
	}

	questions := []Question{}
	answers := []Answer{}

	for _, question := range questionsEntity {
		answersEntity, err := s.answerStore.GetAnswersForQuestion(ctx, question.ID)
		if err != nil {
			return []Question{}, err
		}

		for _, answer := range answersEntity {
			answers = append(answers, Answer{
				answer.ID,
				answer.Body,
				answer.Correct,
				answer.QuestionID,
			})
		}

		questions = append(questions, Question{
			ID:      question.ID,
			Body:    question.Body,
			Answers: answers,
		})
	}

	return questions, err
}
