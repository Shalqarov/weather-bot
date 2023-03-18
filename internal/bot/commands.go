package bot

import (
	"log"

	"github.com/Shalqarov/weather-bot/internal/weather"
	"github.com/Shalqarov/weather-bot/pkg"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const welcomeMsg = "Через команду /w *Название локации* можно узнать текущую погоду в данной локации"

func handleMessages(update tgbotapi.Update, msg *tgbotapi.MessageConfig) error {
	var err error
	users := pkg.UsersMap()
	switch update.Message.Command() {
	case "start":
		msg.Text = welcomeMsg
		log.Println(welcomeMsg)
		users.Set(update.Message.Chat.ID)
	case "w":
		msg.Text, err = weather.CurrentTemperature(update.Message.CommandArguments())
		if err != nil {
			msg.Text = "Неверно указана локация"
		}
	}
	return err
}
