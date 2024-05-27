package telegram

import (
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/zakether/go-interview-bot/internal/cases"
	"github.com/zakether/go-interview-bot/internal/entities"
	"github.com/zakether/go-interview-bot/pkg/logger"
	"go.uber.org/zap"
)

type Bot struct {
	bot             *tgbotapi.BotAPI
	questionService *cases.QuestionService
	sessions        map[int64]entities.Session
	results         map[int64][]entities.Question
	logger          logger.Logger
}

func NewBot(bot *tgbotapi.BotAPI, questionService *cases.QuestionService, logger logger.Logger) *Bot {
	sessions := make(map[int64]entities.Session)
	results := make(map[int64][]entities.Question)
	return &Bot{
		bot:             bot,
		questionService: questionService,
		sessions:        sessions,
		results:         results,
		logger:          logger,
	}
}

func (b *Bot) HandleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.CallbackQuery != nil {
			chatID := update.CallbackQuery.Message.Chat.ID
			userID := int64(update.CallbackQuery.Message.From.ID)
			data := update.CallbackQuery.Data

			if strings.HasPrefix(data, "grade_") {
				b.handleGradeSelection(chatID, userID, data)
			} else {
				session, ok := b.sessions[chatID]
				if !ok {
					b.sendMsg(chatID, "Сессия не найдена. Начните новый тест.")
					continue
				}

				if data == "show_answers" {
					b.handleShowAnswers(update)
				} else {
					b.handleCallbackQuery(chatID, data, &session)
					b.sessions[chatID] = session
				}
			}
		} else if update.Message != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			msgText := update.Message.Text

			switch msgText {
			case "/start":
				b.start(update)
			case "/help":
				b.help(update)
			case "/test":
				b.test(update)
			default:
				b.defaultMsg(update)
			}
		}
	}
}

func (b *Bot) handleGradeSelection(chatID int64, userID int64, data string) {
	var grade string
	switch data {
	case "grade_junior":
		grade = "junior"
	case "grade_middle":
		grade = "middle"
	default:
		b.sendMsg(chatID, "Неверный выбор уровня сложности или типа вопросов.")
		return
	}

	session, err := b.questionService.StartTest(userID, grade)
	if err != nil {
		b.sendMsg(chatID, "Ошибка при начале тестирования.")
		b.logger.Error("Ошибка при начале тестирования", zap.Error(err))
		return
	}

	session.ExpiredAt = time.Now().Add(7 * time.Minute)
	session.CurrentQuestionIndex = 0
	b.sessions[chatID] = session

	b.logger.Info("Отправка первого вопроса")
	b.sendQuestion(chatID, session)
	b.sendRemainingTime(chatID, &session)
}
