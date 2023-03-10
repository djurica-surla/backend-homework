package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/djurica-surla/backend-homework/internal/entity"
	"github.com/djurica-surla/backend-homework/internal/mock/questionOptionStorerMock"
	"github.com/djurica-surla/backend-homework/internal/mock/questionStorerMock"
	"github.com/djurica-surla/backend-homework/internal/service"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

// Service mocks.
type Mocks struct {
	questionStorer       *questionStorerMock.MockQuestionStorer
	questionOptionStorer *questionOptionStorerMock.MockQuestionOptionStorer
}

func createMocks(ctrl *gomock.Controller) Mocks {
	return Mocks{
		questionStorer:       questionStorerMock.NewMockQuestionStorer(ctrl),
		questionOptionStorer: questionOptionStorerMock.NewMockQuestionOptionStorer(ctrl),
	}
}

func initMockService(t *testing.T) (Mocks, *service.QuestionService) {
	ctrl := gomock.NewController(t)

	mocks := createMocks(ctrl)

	svc := service.NewQuestionService(mocks.questionStorer, mocks.questionOptionStorer)

	assert.NotEmpty(t, svc)

	return mocks, svc
}

func TestService_GetQuestions(t *testing.T) {
	t.Run("Should retrieve questions successfuly", func(t *testing.T) {
		ctx := context.Background()
		pageSize := 10
		offset := 0
		mocks, svc := initMockService(t)
		returnQuestions := []entity.Question{
			{
				ID:   1,
				Body: "first-question",
			},
			{
				ID:   2,
				Body: "second-question",
			},
		}
		returnQuestionOptions := []entity.QuestionOption{
			{
				ID:      1,
				Body:    "first-option",
				Correct: false,
			},
			{
				ID:      2,
				Body:    "second-option",
				Correct: false,
			},
			{
				ID:      3,
				Body:    "third-option",
				Correct: true,
			},
		}

		expectedResult := []service.QuestionDTO{
			{
				ID:   1,
				Body: "first-question",
				Options: []service.QuestionOptionDTO{
					{
						ID:      1,
						Body:    "first-option",
						Correct: false,
					},
					{
						ID:      2,
						Body:    "second-option",
						Correct: false,
					},
					{
						ID:      3,
						Body:    "third-option",
						Correct: true,
					},
				},
			},
			{
				ID:   2,
				Body: "second-question",
				Options: []service.QuestionOptionDTO{
					{
						ID:      1,
						Body:    "first-option",
						Correct: false,
					},
					{
						ID:      2,
						Body:    "second-option",
						Correct: false,
					},
					{
						ID:      3,
						Body:    "third-option",
						Correct: true,
					},
				},
			},
		}

		gomock.InOrder(
			mocks.questionStorer.EXPECT().GetQuestions(ctx, pageSize, offset).Return(returnQuestions, nil),
			mocks.questionOptionStorer.EXPECT().GetQuestionOptions(ctx, 1).Return(returnQuestionOptions, nil),
			mocks.questionOptionStorer.EXPECT().GetQuestionOptions(ctx, 2).Return(returnQuestionOptions, nil),
		)

		questions, err := svc.GetQuestions(ctx, pageSize, offset)
		assert.EqualValues(t, expectedResult, questions)
		assert.NoError(t, err)
	})
	t.Run("Should fail because getting questions from database fails", func(t *testing.T) {
		ctx := context.Background()
		pageSize := 10
		offset := 0
		mocks, svc := initMockService(t)
		someErr := errors.New("some-error")

		gomock.InOrder(
			mocks.questionStorer.EXPECT().GetQuestions(ctx, pageSize, offset).Return(nil, someErr),
		)

		questions, err := svc.GetQuestions(ctx, pageSize, offset)
		assert.Nil(t, questions)
		assert.Error(t, err)
	})

	t.Run("Should fail because getting question options from database fails", func(t *testing.T) {
		ctx := context.Background()
		pageSize := 10
		offset := 0
		mocks, svc := initMockService(t)
		someErr := errors.New("some-error")

		returnQuestions := []entity.Question{
			{
				ID:   1,
				Body: "first-question",
			},
		}

		gomock.InOrder(
			mocks.questionStorer.EXPECT().GetQuestions(ctx, pageSize, offset).Return(returnQuestions, nil),
			mocks.questionOptionStorer.EXPECT().GetQuestionOptions(ctx, 1).Return(nil, someErr),
		)

		questions, err := svc.GetQuestions(ctx, pageSize, offset)
		assert.Nil(t, questions)
		assert.Error(t, err)
	})
}

func TestService_GetQuestionByID(t *testing.T) {
	t.Run("Should retrieve question successfuly", func(t *testing.T) {
		ctx := context.Background()
		mocks, svc := initMockService(t)

		returnQuestion := entity.Question{
			ID:   1,
			Body: "first-question",
		}

		returnQuestionOptions := []entity.QuestionOption{
			{
				ID:      1,
				Body:    "first-option",
				Correct: false,
			},
			{
				ID:      2,
				Body:    "second-option",
				Correct: false,
			},
			{
				ID:      3,
				Body:    "third-option",
				Correct: true,
			},
		}

		expectedResult := service.QuestionDTO{
			ID:   1,
			Body: "first-question",
			Options: []service.QuestionOptionDTO{
				{
					ID:      1,
					Body:    "first-option",
					Correct: false,
				},
				{
					ID:      2,
					Body:    "second-option",
					Correct: false,
				},
				{
					ID:      3,
					Body:    "third-option",
					Correct: true,
				},
			},
		}

		gomock.InOrder(
			mocks.questionStorer.EXPECT().GetQuestionByID(ctx, 1).Return(returnQuestion, nil),
			mocks.questionOptionStorer.EXPECT().GetQuestionOptions(ctx, 1).Return(returnQuestionOptions, nil),
		)

		questions, err := svc.GetQuestionByID(ctx, 1)
		assert.EqualValues(t, expectedResult, questions)
		assert.NoError(t, err)
	})

	t.Run("Should fail because getting question from database fails", func(t *testing.T) {
		ctx := context.Background()

		mocks, svc := initMockService(t)
		someErr := errors.New("some-error")

		gomock.InOrder(
			mocks.questionStorer.EXPECT().GetQuestionByID(ctx, 1).Return(entity.Question{}, someErr),
		)

		questions, err := svc.GetQuestionByID(ctx, 1)
		assert.Equal(t, service.QuestionDTO{}, questions)
		assert.Error(t, err)
	})

	t.Run("Should fail because getting question options from database fails", func(t *testing.T) {
		ctx := context.Background()

		mocks, svc := initMockService(t)
		someErr := errors.New("some-error")

		returnQuestion := entity.Question{
			ID:   1,
			Body: "first-question",
		}

		gomock.InOrder(
			mocks.questionStorer.EXPECT().GetQuestionByID(ctx, 1).Return(returnQuestion, nil),
			mocks.questionOptionStorer.EXPECT().GetQuestionOptions(ctx, 1).Return(nil, someErr),
		)

		questions, err := svc.GetQuestionByID(ctx, 1)
		assert.Equal(t, service.QuestionDTO{}, questions)
		assert.Error(t, err)
	})
}

func TestService_CreateQuestion(t *testing.T) {
	t.Run("Should create question successfuly", func(t *testing.T) {
		ctx := context.Background()
		mocks, svc := initMockService(t)

		questionCreationDTO := service.QuestionCreationDTO{
			Body: "first-question",
			Options: []service.QuestionOptionCreationDTO{
				{
					Body:    "first-option",
					Correct: false,
				},
				{
					Body:    "second-option",
					Correct: false,
				},
			},
		}

		optionCreationDTO1 := service.QuestionOptionCreationDTO{
			Body:    "first-option",
			Correct: false,
		}
		optionCreationDTO2 := service.QuestionOptionCreationDTO{
			Body:    "second-option",
			Correct: false,
		}

		storedQuestion := entity.Question{
			ID:   1,
			Body: "first-question",
		}

		storedQuestionOption := []entity.QuestionOption{
			{ID: 1,
				Body:    "first-option",
				Correct: false,
			},
			{
				ID:      2,
				Body:    "second-option",
				Correct: false,
			},
		}

		expectedResult := service.QuestionDTO{
			ID:   1,
			Body: "first-question",
			Options: []service.QuestionOptionDTO{
				{
					ID:      1,
					Body:    "first-option",
					Correct: false,
				},
				{
					ID:      2,
					Body:    "second-option",
					Correct: false,
				},
			},
		}

		gomock.InOrder(
			mocks.questionStorer.EXPECT().CreateQuestion(ctx, questionCreationDTO.Body).Return(1, nil),
			mocks.questionOptionStorer.EXPECT().CreateQuestionOption(ctx, 1, optionCreationDTO1).Return(nil),
			mocks.questionOptionStorer.EXPECT().CreateQuestionOption(ctx, 1, optionCreationDTO2).Return(nil),
			mocks.questionStorer.EXPECT().GetQuestionByID(ctx, 1).Return(storedQuestion, nil),
			mocks.questionOptionStorer.EXPECT().GetQuestionOptions(ctx, 1).Return(storedQuestionOption, nil),
		)

		question, err := svc.CreateQuestion(ctx, questionCreationDTO)
		assert.EqualValues(t, expectedResult, question)
		assert.NoError(t, err)
	})

	t.Run("Should fail because creating question in the database fails", func(t *testing.T) {
		ctx := context.Background()

		mocks, svc := initMockService(t)
		someErr := errors.New("some-error")

		gomock.InOrder(
			mocks.questionStorer.EXPECT().CreateQuestion(ctx, "").Return(0, someErr),
		)

		question, err := svc.CreateQuestion(ctx, service.QuestionCreationDTO{})
		assert.Equal(t, service.QuestionDTO{}, question)
		assert.Error(t, err)
	})

	t.Run("Should fail because creating question options in the database fails", func(t *testing.T) {
		ctx := context.Background()

		mocks, svc := initMockService(t)
		someErr := errors.New("some-error")

		questionCreationDTO := service.QuestionCreationDTO{
			Body: "first-question",
			Options: []service.QuestionOptionCreationDTO{
				{
					Body:    "first-option",
					Correct: false,
				},
			},
		}

		optionCreationDTO1 := service.QuestionOptionCreationDTO{
			Body:    "first-option",
			Correct: false,
		}

		gomock.InOrder(
			mocks.questionStorer.EXPECT().CreateQuestion(ctx, questionCreationDTO.Body).Return(1, nil),
			mocks.questionOptionStorer.EXPECT().CreateQuestionOption(ctx, 1, optionCreationDTO1).Return(someErr),
		)

		question, err := svc.CreateQuestion(ctx, questionCreationDTO)
		assert.Equal(t, service.QuestionDTO{}, question)
		assert.Error(t, err)
	})
}

func TestService_UpdateQuestion(t *testing.T) {
	t.Run("Should update question successfuly", func(t *testing.T) {
		ctx := context.Background()
		mocks, svc := initMockService(t)

		questionCreationDTO := service.QuestionCreationDTO{
			Body: "first-question",
			Options: []service.QuestionOptionCreationDTO{
				{
					Body:    "first-option",
					Correct: false,
				},
				{
					Body:    "second-option",
					Correct: false,
				},
			},
		}

		optionCreationDTO1 := service.QuestionOptionCreationDTO{
			Body:    "first-option",
			Correct: false,
		}
		optionCreationDTO2 := service.QuestionOptionCreationDTO{
			Body:    "second-option",
			Correct: false,
		}

		storedQuestion := entity.Question{
			ID:   1,
			Body: "first-question",
		}

		storedQuestionOption := []entity.QuestionOption{
			{ID: 1,
				Body:    "first-option",
				Correct: false,
			},
			{
				ID:      2,
				Body:    "second-option",
				Correct: false,
			},
		}

		expectedResult := service.QuestionDTO{
			ID:   1,
			Body: "first-question",
			Options: []service.QuestionOptionDTO{
				{
					ID:      1,
					Body:    "first-option",
					Correct: false,
				},
				{
					ID:      2,
					Body:    "second-option",
					Correct: false,
				},
			},
		}

		gomock.InOrder(
			mocks.questionStorer.EXPECT().UpdateQuestion(ctx, 1, questionCreationDTO.Body).Return(1, nil),
			mocks.questionOptionStorer.EXPECT().DeleteQuestionOptions(ctx, 1).Return(nil),
			mocks.questionOptionStorer.EXPECT().CreateQuestionOption(ctx, 1, optionCreationDTO1).Return(nil),
			mocks.questionOptionStorer.EXPECT().CreateQuestionOption(ctx, 1, optionCreationDTO2).Return(nil),
			mocks.questionStorer.EXPECT().GetQuestionByID(ctx, 1).Return(storedQuestion, nil),
			mocks.questionOptionStorer.EXPECT().GetQuestionOptions(ctx, 1).Return(storedQuestionOption, nil),
		)

		question, err := svc.UpdateQuestion(ctx, 1, questionCreationDTO)
		assert.EqualValues(t, expectedResult, question)
		assert.NoError(t, err)
	})

	t.Run("Should return empty dto response because update affects no rows", func(t *testing.T) {
		ctx := context.Background()

		mocks, svc := initMockService(t)

		questionCreationDTO := service.QuestionCreationDTO{
			Body: "first-question",
			Options: []service.QuestionOptionCreationDTO{
				{
					Body:    "first-option",
					Correct: false,
				},
				{
					Body:    "second-option",
					Correct: false,
				},
			},
		}

		gomock.InOrder(
			mocks.questionStorer.EXPECT().UpdateQuestion(ctx, 1, questionCreationDTO.Body).Return(0, nil),
		)

		question, err := svc.UpdateQuestion(ctx, 1, questionCreationDTO)
		assert.Equal(t, service.QuestionDTO{}, question)
		assert.NoError(t, err)
	})

	t.Run("Should return error because deleting previous option fails", func(t *testing.T) {
		ctx := context.Background()

		mocks, svc := initMockService(t)

		someErr := errors.New("some-error")

		questionCreationDTO := service.QuestionCreationDTO{
			Body: "first-question",
			Options: []service.QuestionOptionCreationDTO{
				{
					Body:    "first-option",
					Correct: false,
				},
				{
					Body:    "second-option",
					Correct: false,
				},
			},
		}

		gomock.InOrder(
			mocks.questionStorer.EXPECT().UpdateQuestion(ctx, 1, questionCreationDTO.Body).Return(1, nil),
			mocks.questionOptionStorer.EXPECT().DeleteQuestionOptions(ctx, 1).Return(someErr),
		)

		question, err := svc.UpdateQuestion(ctx, 1, questionCreationDTO)
		assert.Equal(t, service.QuestionDTO{}, question)
		assert.Error(t, err)
	})

	t.Run("Should return error because creating new options fails", func(t *testing.T) {
		ctx := context.Background()

		mocks, svc := initMockService(t)

		someErr := errors.New("some-error")

		questionCreationDTO := service.QuestionCreationDTO{
			Body: "first-question",
			Options: []service.QuestionOptionCreationDTO{
				{
					Body:    "first-option",
					Correct: false,
				},
			},
		}
		optionCreationDTO1 := service.QuestionOptionCreationDTO{
			Body:    "first-option",
			Correct: false,
		}

		gomock.InOrder(
			mocks.questionStorer.EXPECT().UpdateQuestion(ctx, 1, questionCreationDTO.Body).Return(1, nil),
			mocks.questionOptionStorer.EXPECT().DeleteQuestionOptions(ctx, 1).Return(nil),
			mocks.questionOptionStorer.EXPECT().CreateQuestionOption(ctx, 1, optionCreationDTO1).Return(someErr),
		)

		question, err := svc.UpdateQuestion(ctx, 1, questionCreationDTO)
		assert.Equal(t, service.QuestionDTO{}, question)
		assert.Error(t, err)
	})
}

func TestService_DeleteQuestion(t *testing.T) {
	t.Run("Should delete question successfuly", func(t *testing.T) {
		ctx := context.Background()
		mocks, svc := initMockService(t)

		gomock.InOrder(
			mocks.questionStorer.EXPECT().DeleteQuestion(ctx, 1).Return(nil),
		)

		err := svc.DeleteQuestion(ctx, 1)
		assert.NoError(t, err)
	})

	t.Run("Should fail to delete question when database fails", func(t *testing.T) {
		ctx := context.Background()
		mocks, svc := initMockService(t)

		someErr := errors.New("some-error")

		gomock.InOrder(
			mocks.questionStorer.EXPECT().DeleteQuestion(ctx, 1).Return(someErr),
		)

		err := svc.DeleteQuestion(ctx, 1)
		assert.Error(t, err)
	})
}
