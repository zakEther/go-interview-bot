package bot

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	adapters "github.com/zakether/go-interview-bot/internal/adapters/storage"
	"github.com/zakether/go-interview-bot/internal/cases"
	"github.com/zakether/go-interview-bot/internal/infrastructure/delivery/telegram"
	"github.com/zakether/go-interview-bot/pkg/config"
	"github.com/zakether/go-interview-bot/pkg/logger"
	"github.com/zakether/go-interview-bot/pkg/postgres"
	"go.uber.org/zap"
)

func StartBot(path string) error {

	logger, _ := logger.NewLogger()

	const op = "ошибка"
	cfg, err := config.LoadConfig(path)
	if err != nil {
		return err
	}

	db, err := postgres.NewPostgreSQLStorage(*cfg)
	if err != nil {
		return fmt.Errorf("ошибка подключения к базе данных: %w, %s", err, op)
	}

	logger.Info("Initializing telegram bot")

	bot, err := tgbotapi.NewBotAPI(cfg.BotToken)
	if err != nil {
		logger.Fatal("Failed to initialize telegram bot", zap.Error(err))
	}

	logger.Info("Authorized on account" + bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	logger.Info("Bot is running and next step is to get updates")
	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		logger.Fatal("Failed to get updates", zap.Error(err))
	}
	logger.Info("At this point, the bot is running and listening for updates")

	logger.Info("Initializing session storage")
	sessionStorage := adapters.New(db)
	if err != nil {
		return fmt.Errorf("ошибка инициализации хранилища сессий: %w", err)
	}
	logger.Info("Session storage initialized")

	logger.Info("Initializing question service")
	questionService := cases.NewQuestionService(sessionStorage, zap.L())

	logger.Info("Initializing bot instance")
	botInstance := telegram.NewBot(bot, questionService, logger)

	logger.Info("Starting bot instance")

	botInstance.HandleUpdates(updates)

	logger.Info("Bot listening for updates")

	return err

}
