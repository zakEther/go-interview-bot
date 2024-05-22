package telegram

import (
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/rs/zerolog/log"
	"github.com/zakether/go-interview-bot/internal/entities"
	"go.uber.org/zap"
)

func (b *Bot) createKeyboard(question entities.Question, session entities.Session) tgbotapi.InlineKeyboardMarkup {
	var rows [][]tgbotapi.InlineKeyboardButton

	for i, option := range question.QuestionOptions {
		row := tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(option, fmt.Sprintf("answer_%d_%d", question.QuestionID, i)),
		)
		rows = append(rows, row)
	}

	b.logger.Info("Варианты ответов", zap.Strings("options", question.QuestionOptions))

	if session.CurrentQuestionIndex < len(session.Questions)-1 {
		navigationRow := tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Следующий вопрос", "next_question"),
		)
		rows = append(rows, navigationRow)
		b.logger.Info("Добавлена кнопка 'Следующий вопрос'")
	} else {
		finishRow := tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Завершить тест", "submit"),
		)
		rows = append(rows, finishRow)
		b.logger.Info("Добавлена кнопка 'Завершить тест'")
	}

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

	keyboard := b.createKeyboard(currentQuestion, session)

	var msgText strings.Builder
	msgText.WriteString(fmt.Sprintf("# %d: %s\n", session.CurrentQuestionIndex+1, currentQuestion.QuestionText))

	for i, option := range currentQuestion.QuestionOptions {
		msgText.WriteString(fmt.Sprintf("   %d. %s\n", i+1, option))
	}

	msg := tgbotapi.NewMessage(chatID, msgText.String())
	msg.ReplyMarkup = &keyboard

	b.logger.Info("Отправка вопроса", zap.Int("questionIndex", session.CurrentQuestionIndex+1))
	_, err := b.bot.Send(msg)
	if err != nil {
		log.Error().Err(err).Msg("Ошибка при отправке вопроса")
	}
}
