package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/zakether/go-interview-bot/internal/cases"
	"github.com/zakether/go-interview-bot/internal/entities"
	"github.com/zakether/go-interview-bot/pkg/logger"
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
			data := update.CallbackQuery.Data

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
