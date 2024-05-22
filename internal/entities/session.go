package entities

import (
	"time"
)

type Session struct {
	SessionID            string     `json:"session_id"`
	UserID               int64      `json:"user_id"`
	Questions            []Question `json:"questions"`
	CurrentQuestionIndex int        `json:"current_question_index"`
	UserAnswers          []int      `json:"users_answers"`
	ExpiredAt            time.Time  `json:"expired_at"`
}

func NewSession(id string, questions []Question, userID int64) Session {
	return Session{
		SessionID:            id,
		UserID:               userID,
		Questions:            questions,
		CurrentQuestionIndex: 0,
		UserAnswers:          make([]int, len(questions)),
		ExpiredAt:            time.Now().Add(7 * time.Minute),
	}
}
