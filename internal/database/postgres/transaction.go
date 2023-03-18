package postgres

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

func Transaction(ctx context.Context) (*sqlx.Tx, error) {
	tx, err := GetDB().BeginTxx(ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
	})
	if err != nil {
		return nil, err
	}
	return tx, nil
}
