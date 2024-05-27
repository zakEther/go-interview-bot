package cases

import (
	"github.com/zakether/go-interview-bot/internal/entities"
	"go.uber.org/zap"
)

type QuestionService struct {
	storage SessionStorage
	logger  *zap.Logger
}

func NewQuestionService(storage SessionStorage, logger *zap.Logger) *QuestionService {
	return &QuestionService{
		storage: storage,
		logger:  logger,
	}
}

func (qs *QuestionService) StartTest(userID int64, grade string) (entities.Session, error) {
	qs.logger.Info("Создание новой сессии для пользователя", zap.Int64("userID", userID))
	session, err := qs.storage.CreateSession(userID, grade)
	if err != nil {
		qs.logger.Error("Ошибка при создании сессии", zap.Error(err))
		return entities.Session{}, err
	}

	session.CurrentQuestionIndex = 0

	qs.logger.Info("Создана новая сессия", zap.Any("sessionID", session.SessionID))
	return session, err
}

func (qs *QuestionService) GetResult(session *entities.Session, userAnswers []int) (int, []entities.Question, error) {
	err := qs.storage.SubmitUserAnswers(session, userAnswers)
	if err != nil {
		qs.logger.Error("Failed to submit user answers", zap.Error(err))
		return 0, nil, err
	}

	var score int
	for i, answer := range session.UserAnswers {
		if answer == session.Questions[i].Answer {
			score++
		}
	}

	return score, session.Questions, nil
}
