package telegram

import (
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/rs/zerolog/log"
	"github.com/zakether/go-interview-bot/internal/entities"
	"go.uber.org/zap"
)

func (b *Bot) createKeyboard(question entities.Question) tgbotapi.InlineKeyboardMarkup {
	var rows [][]tgbotapi.InlineKeyboardButton

	for i, option := range question.QuestionOptions {
		row := tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(option, fmt.Sprintf("answer_%d_%d", question.QuestionID, i)),
		)
		rows = append(rows, row)
	}

	b.logger.Info("Варианты ответов", zap.Strings("options", question.QuestionOptions))

	keyboard := tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: rows,
	}

	return keyboard
}

func (b *Bot) sendQuestion(chatID int64, session entities.Session) {
	if len(session.Questions) == 0 {
		log.Error().Msg("Нет вопросов для тестирования")
		b.sendMsg(chatID, "Не удалось загрузить вопросы для теста.")
		return
	}

	currentQuestion := session.Questions[session.CurrentQuestionIndex]

	keyboard := b.createKeyboard(currentQuestion)

	var msgText strings.Builder
	msgText.WriteString(fmt.Sprintf("# %d: %s\n", session.CurrentQuestionIndex+1, currentQuestion.QuestionText))

	msg := tgbotapi.NewMessage(chatID, msgText.String())
	msg.ParseMode = tgbotapi.ModeMarkdown
	msg.ReplyMarkup = &keyboard

	b.logger.Info("Отправка вопроса", zap.Int("questionIndex", session.CurrentQuestionIndex+1))
	_, err := b.bot.Send(msg)
	if err != nil {
		log.Error().Err(err).Msg("Ошибка при отправке вопроса")
	}
}
