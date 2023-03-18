package stats

import (
	"context"
	"fmt"
	"time"

	"github.com/Shalqarov/weather-bot/internal/database/postgres/repository/stat"
	"github.com/Shalqarov/weather-bot/tools"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Handler(update tgbotapi.Update, msg *tgbotapi.MessageConfig) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	count, err := stat.CountByUser(ctx, update.Message.From.ID)
	if err != nil {
		msg.Text = tools.ErrorMsg
		return err
	}

	ctx, cancel = context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	statistic, err := stat.EarliestByUser(ctx, update.Message.From.ID)
	if err != nil {
		msg.Text = tools.ErrorMsg
		return err
	}
	msg.Text = fmt.Sprintf(
		"Ваша статистика:\n"+
			"Количество запросов: %d\n"+
			"Самая ранняя запись:\n%s\nОтправлено: GMT %v\n",
		count, statistic.Message, statistic.CreatedAt.Format("01-02-2006 15:04 MST"),
	)
	return nil
}
