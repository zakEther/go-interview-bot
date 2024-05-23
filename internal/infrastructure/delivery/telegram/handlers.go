package telegram

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/rs/zerolog/log"
	"github.com/zakether/go-interview-bot/internal/entities"
	"go.uber.org/zap"
)

func (b *Bot) start(request tgbotapi.Update) {
	chatID := request.Message.Chat.ID
	message := "У Вас будет 7 минут для того, чтобы ответить на 15 вопросов\n" +
		"Ознакомьтесь с правилами - /help\n" +
		"Нажмите /test для начала тестирования"
	b.sendMsg(chatID, message)
}

func (b *Bot) help(request tgbotapi.Update) {
	chatID := request.Message.Chat.ID
	message := "🤖Я бот для проведения тестирования по языку Go\n" +
		"У Вас будет 7 минут для сдачи всех вопросов\n" +
		"❓<b>Как пройти тест?</b>\n" +
		"1. Нажимаете на /test\n" +
		"2. Выбираете ответ\n\n" +
		"3. Если нашли ошибку - пишите @zakether\n\n" +
		"🌟Удачи"

	msg := tgbotapi.NewMessage(chatID, message)
	msg.ParseMode = "HTML"
	b.bot.Send(msg)
}

func (b *Bot) test(request tgbotapi.Update) {
	chatID := request.Message.Chat.ID
	userID := int64(request.Message.From.ID)

	session, err := b.questionService.StartTest(userID)
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

func (b *Bot) defaultMsg(request tgbotapi.Update) {
	msg := "Не могу распознать команду. Кликни /help для просмотра команд"
	b.sendMsg(request.Message.Chat.ID, msg)
}

func (b *Bot) sendMsg(chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)
	_, err := b.bot.Send(msg)
	if err != nil {
		log.Error().Err(err).Msg("Ошибка при отправке сообщения")
	}
}

func (b *Bot) handleCallbackQuery(chatID int64, data string, session *entities.Session) {
	switch {
	case data == "submit":
		b.logger.Info("Завершение теста", zap.Int64("userID", session.UserID))
		score, _, err := b.questionService.GetResult(session, session.UserAnswers)
		if err != nil {
			b.logger.Error("Ошибка при завершении теста", zap.Error(err))
			b.sendMsg(chatID, "Произошла ошибка при завершении теста.")
			return
		}
		resultMsg := fmt.Sprintf("Тест завершен. Ваш результат: %d", score)
		b.sendMsg(chatID, resultMsg)

		button := tgbotapi.NewInlineKeyboardButtonData("Посмотреть ответы", "show_answers")
		keyboard := tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(button))
		msg := tgbotapi.NewMessage(chatID, "Вы можете просмотреть правильные ответы, нажав на кнопку ниже.")
		msg.ReplyMarkup = keyboard
		b.bot.Send(msg)

		b.sessions[chatID] = *session

	default:
		if strings.HasPrefix(data, "answer_") {
			parts := strings.Split(data, "_")
			if len(parts) == 3 {
				questionID, err1 := strconv.Atoi(parts[1])
				answerIndex, err2 := strconv.Atoi(parts[2])
				if err1 == nil && err2 == nil {
					b.logger.Info("Ответ получен", zap.Int("questionID", questionID), zap.Int("answerIndex", answerIndex))
					session.UserAnswers[session.CurrentQuestionIndex] = answerIndex
					session.CurrentQuestionIndex++
					b.sessions[chatID] = *session

					if session.CurrentQuestionIndex < len(session.Questions) {
						b.sendQuestion(chatID, *session)
						b.sendRemainingTime(chatID, session)
					} else {
						b.handleCallbackQuery(chatID, "submit", session)
					}
				}
			}
		}
	}
}

func (b *Bot) handleShowAnswers(update tgbotapi.Update) {
	chatID := update.CallbackQuery.Message.Chat.ID

	session, ok := b.sessions[chatID]
	if !ok {
		b.sendMsg(chatID, "Сессия не найдена. Начните новый тест.")
		return
	}

	var results []string
	for i, answerIndex := range session.UserAnswers {
		correctAnswerIndex := session.Questions[i].GetAnswer()
		if answerIndex != correctAnswerIndex {
			question := session.Questions[i]
			result := fmt.Sprintf(
				"***Вопрос:*** %s\n***Ваш ответ:*** %s\n***Правильный ответ:*** %s\n***Объяснение:*** %s\n",
				question.GetText(),
				question.GetQuestionOptions()[answerIndex],
				question.GetQuestionOptions()[correctAnswerIndex],
				question.Explanation)
			results = append(results, result)
		}
	}

	if len(results) == 0 {
		b.sendMsg(chatID, "Поздравляем! Все ответы правильные.")
	} else {
		for _, result := range results {
			msg := tgbotapi.NewMessage(chatID, result)
			msg.ParseMode = tgbotapi.ModeMarkdown
			b.bot.Send(msg)
		}
	}
}

func (b *Bot) sendRemainingTime(chatID int64, session *entities.Session) {
	timeLeft := time.Until(session.ExpiredAt)
	minutes := int(timeLeft.Minutes())
	seconds := int(timeLeft.Seconds()) % 60

	timeLeftMsg := fmt.Sprintf("До окончания теста: %dм %dс", minutes, seconds)
	b.sendMsg(chatID, timeLeftMsg)
}
