package register

import (
	"context"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/Shalqarov/weather-bot/internal/database/postgres"
	"github.com/Shalqarov/weather-bot/internal/database/postgres/repository/user"
	"github.com/Shalqarov/weather-bot/tools"
)

const (
	welcomeMsg = "Через команду /w *Название локации* можно узнать текущую погоду в данной локации"
)

func Handler(update tgbotapi.Update, msg *tgbotapi.MessageConfig) error {
	msg.Text = welcomeMsg
	msg.ReplyToMessageID = update.Message.MessageID

	u := user.User{
		UserID:    update.Message.From.ID,
		CreatedAt: time.Now(),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	tx, err := postgres.Transaction(ctx)
	if err != nil {
		msg.Text = tools.ErrorMsg
		return err
	}
	has, err := user.Has(ctx, u.UserID)
	if err != nil {
		msg.Text = tools.ErrorMsg
		return err
	}
	if has {
		msg.Text = "Вы уже зарегистрированы"
		return nil
	}

	ctx, cancel = context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err = u.Add(ctx, tx); err != nil {
		msg.Text = tools.ErrorMsg
		return err
	}
	if err = tx.Commit(); err != nil {
		msg.Text = tools.ErrorMsg
		return err
	}
	return nil
}
