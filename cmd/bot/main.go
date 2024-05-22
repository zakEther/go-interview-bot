package main

import (
	"log"

	"github.com/zakether/go-interview-bot/internal/bot"
)

func main() {
	path := "config/config.yml"
	if err := bot.StartBot(path); err != nil {
		log.Fatalf("Ошибка при запуске бота: %v", err)
	}

}
