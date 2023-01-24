package entity

// Represents question.
type Question struct {
	ID   int    `json:"id"`
	Body string `json:"body"`
}

// Represents options for question.
type QuestionOption struct {
	ID         int    `json:"id"`
	Body       string `json:"body"`
	Correct    int    `json:"correct"`
	QuestionID int    `json:"question_id"`
}
