package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/djurica-surla/backend-homework/internal/helpers"
	"github.com/djurica-surla/backend-homework/internal/service"
	"github.com/gorilla/mux"
)

// RegisterRoutes links routes with the handler.
func (h *QuestionHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/questions", h.GetQuestions()).Methods(http.MethodGet)
	router.HandleFunc("/questions", h.CreateQuestion()).Methods(http.MethodPost)
	router.HandleFunc("/questions/{id}", h.UpdateQuestion()).Methods(http.MethodPut)
	router.HandleFunc("/questions/{id}", h.DeleteQuestion()).Methods(http.MethodDelete)
}

// QuestionServicer represents necessary question service implementation for question handler.
type QuestionServicer interface {
	GetQuestions(ctx context.Context) ([]service.QuestionDTO, error)
	CreateQuestion(ctx context.Context, questionCreation service.QuestionCreationDTO) (service.QuestionDTO, error)
	UpdateQuestion(ctx context.Context, questionID int, questionCreation service.QuestionCreationDTO) (service.QuestionDTO, error)
	DeleteQuestion(ctx context.Context, questionID int) error
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

// GetQuestions handles retrieveing questions.
func (h *QuestionHandler) GetQuestions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := h.questionService.GetQuestions(r.Context())
		if err != nil {
			h.encodeErrorWithStatus500(err, w)
			return
		}

		json.NewEncoder(w).Encode(res)
	}
}

// CreateQuestion handles creation of questions.
func (h *QuestionHandler) CreateQuestion() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		questionCreationDTO := service.QuestionCreationDTO{}

		err := json.NewDecoder(r.Body).Decode(&questionCreationDTO)
		if err != nil {
			h.encodeErrorWithStatus500(err, w)
			return
		}

		err = helpers.ValidateStruct(questionCreationDTO)
		if err != nil {
			h.encodeErrorWithStatus500(err, w)
			return
		}

		res, err := h.questionService.CreateQuestion(r.Context(), questionCreationDTO)
		if err != nil {
			h.encodeErrorWithStatus500(err, w)
			return
		}

		json.NewEncoder(w).Encode(res)
	}
}

// UpdateQuestion handles updating of questions.
func (h *QuestionHandler) UpdateQuestion() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		questionIDNum, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
		if err != nil {
			h.encodeErrorWithStatus500(err, w)
		}
		questionID := int(questionIDNum)

		questionCreationDTO := service.QuestionCreationDTO{}

		err = json.NewDecoder(r.Body).Decode(&questionCreationDTO)
		if err != nil {
			h.encodeErrorWithStatus500(err, w)
			return
		}

		err = helpers.ValidateStruct(questionCreationDTO)
		if err != nil {
			h.encodeErrorWithStatus500(err, w)
			return
		}

		res, err := h.questionService.UpdateQuestion(r.Context(), questionID, questionCreationDTO)
		if err != nil {
			h.encodeErrorWithStatus500(err, w)
			return
		}

		json.NewEncoder(w).Encode(res)
	}
}

// DeleteQuestion handles deleting of questions.
func (h *QuestionHandler) DeleteQuestion() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		questionIDNum, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
		if err != nil {
			h.encodeErrorWithStatus500(err, w)
		}
		questionID := int(questionIDNum)

		err = h.questionService.DeleteQuestion(r.Context(), questionID)
		if err != nil {
			h.encodeErrorWithStatus500(err, w)
			return
		}

		json.NewEncoder(w).Encode("successfuly deleted question")
	}
}

func (h *QuestionHandler) encodeErrorWithStatus500(err error, w http.ResponseWriter) {
	errorResponse := fmt.Sprintf("error: %s", err.Error())
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(errorResponse)
}
