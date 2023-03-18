package app

import (
	"log"

	"github.com/Shalqarov/weather-bot/internal/bot"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Run() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetBot().GetUpdatesChan(u)
	err := bot.HandleUpdates(updates)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
