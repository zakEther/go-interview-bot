package cases

import "github.com/zakether/go-interview-bot/internal/entities"

type SessionStorage interface {
	CreateSession(userID int64) (entities.Session, error)
	GetRandomQuestions(numQuestions int) ([]entities.Question, error)
	SubmitUserAnswers(session *entities.Session, answers []int) error
}
