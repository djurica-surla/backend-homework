package service

// Represents options for question.
type QuestionOption struct {
	ID      int    `json:"id"`
	Body    string `json:"body"`
	Correct int    `json:"correct"`
}

// Represents question.
type Question struct {
	ID      int              `json:"id"`
	Body    string           `json:"body"`
	Options []QuestionOption `json:"options"`
}
