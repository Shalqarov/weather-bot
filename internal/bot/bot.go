package bot

import (
	"log"
	"sync"

	"github.com/Shalqarov/weather-bot/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	bot   *tgbotapi.BotAPI
	bOnce sync.Once
)

func GetBot() *tgbotapi.BotAPI {
	bOnce.Do(func() {
		var err error
		bot, err = tgbotapi.NewBotAPI(config.GetConfig().TelegramAPI)
		if err != nil {
			log.Fatalln(err.Error())
		}
		log.Println("Bot initialized")
	})
	return bot
}

func HandleUpdates(updates tgbotapi.UpdatesChannel) error {
	var err error

	for update := range updates {

		if update.Message == nil {
			continue
		}
		if !update.Message.IsCommand() {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		err = handleMessages(update, &msg)
		if err != nil {
			log.Println(err.Error())
		}
		_, err = GetBot().Send(msg)
		if err != nil {
			return err
		}
	}
	return nil
}
