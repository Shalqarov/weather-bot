package stats

import (
	"context"
	"time"

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
