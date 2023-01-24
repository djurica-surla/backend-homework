package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/djurica-surla/backend-homework/internal/service"
	"github.com/gorilla/mux"
)

// RegisterRoutes links routes with the handler.
func (h *QuestionHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/questions", h.GetQuestions()).Methods(http.MethodGet)
}

// QuestionServicer represents necessary question service implementation for question handler.
type QuestionServicer interface {
	GetQuestions(ctx context.Context) ([]service.Question, error)
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
