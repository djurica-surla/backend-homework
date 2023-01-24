package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/djurica-surla/backend-homework/internal/service"
	"github.com/gorilla/mux"
)

// RegisterRoutes links routes with the handler
func (h *QuestionHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/questions", h.GetQuestions()).Methods(http.MethodGet)
}

// QuestionServicer represents necessary question service implementation for question handler.
type QuestionServicer interface {
	GetQuestions(ctx context.Context) ([]service.Question, error)
}

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
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(err)
			return
		}

		json.NewEncoder(w).Encode(questions)
	}

}
