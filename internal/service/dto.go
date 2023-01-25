package service

// Question option dto used for response.
type QuestionOptionDTO struct {
	ID      int    `json:"id"`
	Body    string `json:"body"`
	Correct bool   `json:"correct"`
}

// Question dto used for response.
type QuestionDTO struct {
	ID      int                 `json:"id"`
	Body    string              `json:"body"`
	Options []QuestionOptionDTO `json:"options"`
}

// Question option dto used for create and update request.
type QuestionOptionCreationDTO struct {
	Body    string `json:"body" validate:"required"`
	Correct bool   `json:"correct"`
}

// Question dto used for create and update requests.
type QuestionCreationDTO struct {
	Body    string                      `json:"body" validate:"required"`
	Options []QuestionOptionCreationDTO `json:"options" validate:"dive,required"`
}
