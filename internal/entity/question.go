package entity

// Represents question.
type Question struct {
	ID   int
	Body string
}

// Represents options for question.
type QuestionOption struct {
	ID         int
	Body       string
	Correct    bool
	QuestionID int
}
