package service

// Represents options for question.
type QuestionOptionDTO struct {
	ID      int    `json:"id"`
	Body    string `json:"body"`
	Correct bool   `json:"correct"`
}

// Represents question.
type QuestionDTO struct {
	ID      int                 `json:"id"`
	Body    string              `json:"body"`
	Options []QuestionOptionDTO `json:"options"`
}

// Represents question optioon entity for question creation.
type QuestionOptionCreationDTO struct {
	Body    string `json:"body" validate:"required"`
	Correct bool   `json:"correct" validate:"required"`
}

// Represents entity used in question creation.
type QuestionCreationDTO struct {
	Body    string                      `json:"body" validate:"required"`
	Options []QuestionOptionCreationDTO `json:"options" validate:"dive,required"`
}
