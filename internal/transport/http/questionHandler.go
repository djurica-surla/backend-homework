package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/djurica-surla/backend-homework/internal/helpers"
	"github.com/djurica-surla/backend-homework/internal/service"
	"github.com/gorilla/mux"
)

// RegisterRoutes links routes with the handler.
func (h *QuestionHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/questions", h.GetQuestions()).Methods(http.MethodGet)
	router.HandleFunc("/questions", h.CreateQuestion()).Methods(http.MethodPost)
}

// QuestionServicer represents necessary question service implementation for question handler.
type QuestionServicer interface {
	GetQuestions(ctx context.Context) ([]service.QuestionDTO, error)
	CreateQuestion(ctx context.Context, questionCreation service.QuestionCreationDTO) error
}

// QuestionHandler handles http requests for questions.
type QuestionHandler struct {
	questionService QuestionServicer
}

// NewQuestionHandler creates a new instance of question handler.
func NewQuestionHandler(questionService QuestionServicer) *QuestionHandler {
	return &QuestionHandler{
		questionService: questionService,
	}
}

// GetQuestions retrieves questions from the questions service.
func (h *QuestionHandler) GetQuestions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		questions, err := h.questionService.GetQuestions(r.Context())
		if err != nil {
			errorResponse := fmt.Sprintf("error: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(errorResponse)
			return
		}

		json.NewEncoder(w).Encode(questions)
	}
}

// CreateQuestion handles creation of questions.
func (h *QuestionHandler) CreateQuestion() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		questionCreationDTO := service.QuestionCreationDTO{}

		err := json.NewDecoder(r.Body).Decode(&questionCreationDTO)
		if err != nil {
			errorResponse := fmt.Sprintf("error: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(errorResponse)
			return
		}

		err = helpers.ValidateStruct(questionCreationDTO)
		if err != nil {
			errorResponse := fmt.Sprintf("error: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(errorResponse)
			return
		}

		err = h.questionService.CreateQuestion(r.Context(), questionCreationDTO)
		if err != nil {
			errorResponse := fmt.Sprintf("error: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(errorResponse)
			return
		}

		json.NewEncoder(w).Encode("successfully created question!")
	}
}
