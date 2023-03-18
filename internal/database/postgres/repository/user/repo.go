package user

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
)

type User struct {
	UserID    int64     `db:"user_id"`
	CreatedAt time.Time `db:"created_at"`
}

func (u *User) Add(ctx context.Context, tx *sqlx.Tx) error {
	query := `
		INSERT INTO "user" (
		    "user_id",
		    "created_at"
		) VALUES (:user_id, :created_at);`
	_, err := tx.NamedExecContext(ctx, query, u)
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	return nil
}
