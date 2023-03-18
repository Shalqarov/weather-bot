package climate

import (
	"context"
	"time"

	"github.com/Shalqarov/weather-bot/internal/database/postgres"
	"github.com/Shalqarov/weather-bot/internal/database/postgres/repository/stat"
	"github.com/Shalqarov/weather-bot/internal/database/postgres/repository/user"
	"github.com/Shalqarov/weather-bot/internal/weather"
	"github.com/Shalqarov/weather-bot/tools"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	uuid "github.com/satori/go.uuid"
)

func Handler(update tgbotapi.Update, msg *tgbotapi.MessageConfig) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	has, err := user.Has(ctx, update.Message.From.ID)
	if err != nil {
		return err
	}
	if !has {
		msg.Text = "Для ведения статистики запросов, нужно зарегистрироваться через команду /start"
		return nil
	}
	forecast, err := weather.CurrentTemperature(update.Message.CommandArguments())
	if err != nil {
		msg.Text = "Неверно указана локация"
		return err
	}
	msg.Text = forecast

	s := stat.Stat{
		ID:        uuid.NewV4().String(),
		UserID:    update.Message.From.ID,
		Message:   forecast,
		CreatedAt: time.Now(),
	}

	tx, err := postgres.Transaction(ctx)
	if err != nil {
		msg.Text = tools.ErrorMsg
		return err
	}
	ctx, cancel = context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err = s.Add(ctx, tx); err != nil {
		msg.Text = tools.ErrorMsg
		return err
	}
	if err = tx.Commit(); err != nil {
		msg.Text = tools.ErrorMsg
		return err
	}
	return nil
}
