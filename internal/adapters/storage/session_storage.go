package adapters

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/chanxuehong/sid"
	"github.com/lib/pq"
	"github.com/zakether/go-interview-bot/internal/entities"
	"github.com/zakether/go-interview-bot/pkg/postgres"
)

type PGStorage struct {
	DB *postgres.Postgres
}

func New(pg *postgres.Postgres) *PGStorage {
	return &PGStorage{
		DB: pg,
	}
}

func (s *PGStorage) CreateQuestion(question *entities.Question) error {
	_, err := s.DB.Pool.Exec(context.Background(),
		"INSERT INTO questions (question_id, question_text, question_options, answer, explanation) VALUES ($1, $2, $3, $4, $5)",
		question.QuestionID, question.QuestionText, pq.Array(question.QuestionOptions), question.Answer, question.Explanation)
	if err != nil {
		return fmt.Errorf("failed to create question: %w", err)
	}

	return nil
}

func (s *PGStorage) GetRandomQuestions(numQuestions int) ([]entities.Question, error) {
	query := `
		SELECT question_id, question_text, question_options, answer, explanation
		FROM questions
		ORDER BY random()
		LIMIT $1
	`

	rows, err := s.DB.Pool.Query(context.Background(), query, numQuestions)
	if err != nil {
		return nil, fmt.Errorf("failed to query random questions: %v", err)
	}
	defer rows.Close()

	var questions []entities.Question
	for rows.Next() {
		var q entities.Question
		var optionsJSON []byte
		if err := rows.Scan(&q.QuestionID, &q.QuestionText, &optionsJSON, &q.Answer, &q.Explanation); err != nil {
			return nil, fmt.Errorf("failed to scan question row: %v", err)
		}

		err = json.Unmarshal(optionsJSON, &q.QuestionOptions)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal question options: %v", err)
		}

		questions = append(questions, q)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error in row iteration: %v", err)
	}

	return questions, err
}

func (s *PGStorage) CreateSession(userID int64) (entities.Session, error) {
	randomQuestions, err := s.GetRandomQuestions(15)
	if err != nil {
		return entities.Session{}, fmt.Errorf("failed to get random questions: %v", err)
	}

	sessionID := sid.New()
	session := entities.Session{
		SessionID:            sessionID,
		UserID:               userID,
		Questions:            randomQuestions,
		UserAnswers:          make([]int, len(randomQuestions)),
		CurrentQuestionIndex: 0,
	}

	randomQuestionsJSON, err := json.Marshal(session.Questions)
	if err != nil {
		return entities.Session{}, fmt.Errorf("failed to marshal questions: %w", err)
	}

	_, err = s.DB.Pool.Exec(context.Background(),
		"INSERT INTO sessions (session_id, user_id, questions, user_answers, expired_at) VALUES ($1, $2, $3, $4, $5)",
		session.SessionID, session.UserID, randomQuestionsJSON, nil, session.ExpiredAt)

	if err != nil {
		return entities.Session{}, fmt.Errorf("failed to create session: %w", err)
	}

	return session, nil
}

func (s *PGStorage) SubmitUserAnswers(session *entities.Session, answers []int) error {
	session.UserAnswers = answers

	_, err := s.DB.Pool.Exec(context.Background(),
		"UPDATE sessions SET expired_at = $1, user_answers = $2 WHERE session_id = $3",
		session.ExpiredAt, pq.Array(session.UserAnswers), session.SessionID)
	if err != nil {
		return fmt.Errorf("failed to update session score: %w", err)
	}

	return nil
}
