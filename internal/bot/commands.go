package bot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/Shalqarov/weather-bot/internal/weather"
	"github.com/Shalqarov/weather-bot/pkg"
)

const welcomeMsg = "Через команду /w *Название локации* можно узнать текущую погоду в данной локации"

func handleMessages(update tgbotapi.Update, msg *tgbotapi.MessageConfig) (err error) {
	users := pkg.UsersMap()
	switch update.Message.Command() {
	case "start":
		// TODO register handler
		msg.Text = welcomeMsg
		log.Printf("%s has been registered", update.Message.Chat.UserName)
		users.Set(update.Message.Chat.ID)
	case "w":
		msg.Text, err = weather.CurrentTemperature(update.Message.CommandArguments())
		if err != nil {
			msg.Text = "Неверно указана локация"
		}
	case "stat":
		// TODO Когда был первый запрос, сколько запросов было
	}
	return err
}
