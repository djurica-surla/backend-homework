package gen

import (
	// Necessary for proper functioning of mockgen.
	_ "github.com/golang/mock/mockgen/model"
)

//go:generate mockgen -destination=internal/mock/questionStorerMock/questionStorerMock.go -package=questionStorerMock github.com/djurica-surla/backend-homework/internal/service QuestionStorer
//go:generate mockgen -destination=internal/mock/questionOptionStorerMock/questionOptionStorerMock.go -package=questionOptionStorerMock github.com/djurica-surla/backend-homework/internal/service QuestionOptionStorer
