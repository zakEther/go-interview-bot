package cases

import "github.com/zakether/go-interview-bot/internal/entities"

type SessionStorage interface {
	CreateSession(userID int64, grade string) (entities.Session, error)
	GetRandomQuestions(numQuestions int, grade string) ([]entities.Question, error)
	SubmitUserAnswers(session *entities.Session, answers []int) error
}
