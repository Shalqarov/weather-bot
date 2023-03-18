package bot

import (
	"github.com/Shalqarov/weather-bot/internal/handlers/climate"
	"github.com/Shalqarov/weather-bot/internal/handlers/register"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func handleMessages(update tgbotapi.Update, msg *tgbotapi.MessageConfig) (err error) {
	switch update.Message.Command() {
	case "start":
		err = register.Handler(update, msg)
	case "w":
		err = climate.Handler(update, msg)
	case "stat":
		// TODO Когда был первый запрос, сколько запросов было
	}
	return err
}
