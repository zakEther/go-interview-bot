package adapters

// import (
// 	"testing"
// 	"time"

// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/require"
// 	"github.com/zakether/go-interview-bot/internal/entities"
// 	"github.com/zakether/go-interview-bot/pkg/config"
// 	"github.com/zakether/go-interview-bot/pkg/postgres"
// 	"go.uber.org/zap"
// )

// func setupTestPostgreSQLStorage(t *testing.T) (*PGStorage, func()) {
// 	logger, _ := zap.NewDevelopment()

// 	// Создаем конфигурацию для подключения к PostgreSQL
// 	cfg := config.StorageConfig{
// 		// Настройки подключения к базе данных, например:
// 		Host:     "localhost",
// 		Port:     5432,
// 		User:     "username",
// 		Password: "password",
// 		Database: "database_name",
// 		SSLMode:  "disable",
// 	}

// 	// Создаем экземпляр Postgres и обрабатываем возможную ошибку
// 	pg, err := postgres.NewPostgreSQLStorage(cfg)
// 	require.NoError(t, err, "Failed to create PostgreSQL storage")
// 	storage := New(pg)

// 	// Возвращаем хранилище и функцию для закрытия соединения с PostgreSQL
// 	return storage, func() {
// 		pg.Close()
// 	}
// }

// func TestCreateSession(t *testing.T) {
// 	storage, close := setupTestPostgreSQLStorage(t)
// 	defer close()

// 	userID := int64(1)
// 	session, err := storage.CreateSession(userID)
// 	require.NoError(t, err, "CreateSession should not return an error")
// 	assert.NotNil(t, session, "Session should not be nil")
// 	assert.Equal(t, userID, session.UserID, "UserID should match")
// 	assert.Len(t, session.Questions, 10, "Number of questions should match")
// 	assert.NotEmpty(t, session.SessionID, "SessionID should not be empty")
// 	assert.WithinDuration(t, time.Now(), session.ExpiredAt, 1*time.Hour, "ExpiredAt should be within 1 hour from now")
// }

// func TestSubmitUserAnswers(t *testing.T) {
// 	storage, close := setupTestPostgreSQLStorage(t)
// 	defer close()

// 	// Создаем исходную сессию
// 	userID := int64(1)
// 	session, err := storage.CreateSession(userID)
// 	require.NoError(t, err, "CreateSession should not return an error")

// 	// Устанавливаем вопросы с ответами пользователя
// 	userAnswers := make([]int, len(session.Questions))
// 	for i := range userAnswers {
// 		userAnswers[i] = i % len(session.Questions[i].QuestionOptions) // Устанавливаем произвольные ответы
// 	}

// 	err = storage.SubmitUserAnswers(&session, userAnswers)
// 	require.NoError(t, err, "SubmitUserAnswers should not return an error")

// 	// Проверяем, что ответы пользователя сохранены в сессии
// 	assert.Equal(t, userAnswers, session.UserAnswers, "UserAnswers should match")
// }

// func Test_CreateSession(t *testing.T) {
// 	storage, close := setupTestPostgreSQLStorage(t)
// 	defer close()

// 	session := entities.Session{
// 		SessionID: 3,
// 		StartTime: time.Now(),
// 	}

// 	numQuestions := 10
// 	err := storage.CreateSession(&session)
// 	assert.NoError(t, err, "CreateSession should not return an error")
// 	assert.NotNil(t, session, "Session should not be nil")
// 	assert.Len(t, session.Questions, numQuestions, "Number of questions should match")
// 	assert.NotEmpty(t, session.StartTime, "StartTime should not be empty")

// }

// func Test_SaveSession(t *testing.T) {
// 	storage, close := setupTestPostgreSQLStorage(t)
// 	defer close()

// 	// Создание тестовой сессии
// 	session := entities.Session{
// 		SessionID:      2,
// 		Questions:      nil,
// 		StartTime:      time.Now(),
// 		EndTime:        time.Now().Add(1 * time.Hour),
// 		UserAnswers:    []int{},
// 		CorrectAnswers: nil,
// 	}

// 	err := storage.SaveSession(&session)
// 	require.NoError(t, err, "Failed to save session")

// }

// func Test_SessionByID(t *testing.T) {
// 	storage, close := setupTestPostgreSQLStorage(t)
// 	defer close()

// 	testSessions := entities.Session{
// 		SessionID: 3,
// 		StartTime: time.Now(),
// 		EndTime:   time.Now().Add(1 * time.Hour),
// 	}

// 	session, err := storage.SessionByID(testSessions.SessionID)
// 	require.NoError(t, err, "Failed to retrieve session by ID")

// 	// Проверка, что сессия была получена успешно
// 	assert.NotNil(t, session, "Session should not be nil")
// 	assert.Equal(t, testSessions.SessionID, session.SessionID, "Unexpected session ID")
// }

// func TestQuestionStorage(t *testing.T) {
// 	storage, close := setupTestPostgreSQLStorage(t)
// 	defer close()

// 	question := entities.Question{
// 		QuestionID:      10,
// 		QuestionText:    "What is 2 + 2 + 2 + 3?",
// 		QuestionOptions: []string{"3", "4", "5", "9"},
// 		Answer:          3,
// 	}

// 	err := storage.CreateQuestion(&question)
// 	assert.NoError(t, err, "Failed to create question")
// }
