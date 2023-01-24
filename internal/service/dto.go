package service

// Represents answer.
type Answer struct {
	ID         int    `json:"id"`
	Body       string `json:"body"`
	Correct    int    `json:"correct"`
	QuestionID int    `json:"question_id"`
}

// Represents question.
type Question struct {
	ID      int      `json:"id"`
	Body    string   `json:"body"`
	Answers []Answer `json:"answers"`
}
