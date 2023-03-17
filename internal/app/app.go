package app

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Run() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

}
