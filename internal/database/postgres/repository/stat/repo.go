package stats

import (
	"context"
	"time"

	"github.com/Shalqarov/weather-bot/internal/database/postgres"
	"github.com/jmoiron/sqlx"
)

type Stats struct {
	ID        string    `db:"id"`
	UserID    int64     `db:"user_id"`
	Message   string    `db:"message"`
	CreatedAt time.Time `db:"created_at"`
}

func (s *Stats) Add(ctx context.Context, tx *sqlx.Tx) error {
	query := `
		INSERT INTO "stats" (
			"id",
			"user_id",
			"message",
			"created_at"
		) VALUES (:id, :user_id, :message, :created_at);`
	_, err := tx.NamedExecContext(ctx, query, s)
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	return nil
}

func CountByUser(ctx context.Context, id int64) (count int64, err error) {
	query := `
		SELECT COUNT(*)
		FROM "stats" WHERE "user_id" = $1
`
	err = postgres.GetDB().GetContext(ctx, &count, query, id)
	return
}

func (s *Stats) EarliestByUser(ctx context.Context, id int64) (*Stats, error) {
	stat := new(Stats)
	query := `
		SELECT * FROM stats WHERE user_id = $1 ORDER BY created_at LIMIT 1;
`
	err := postgres.GetDB().GetContext(ctx, &stat, query, id)
	if err != nil {
		return nil, err
	}
	return stat, nil
}
