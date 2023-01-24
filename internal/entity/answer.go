package entity

// Represents answer.
type Answer struct {
	ID         int    `json:"id"`
	Body       string `json:"body"`
	Correct    int    `json:"correct"`
	QuestionID int    `json:"question_id"`
}
