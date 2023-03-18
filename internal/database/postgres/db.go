package postgres

import (
	"log"
	"sync"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"

	"github.com/Shalqarov/weather-bot/config"
)

const maxConns = 5

var (
	db     *sqlx.DB
	onceDb sync.Once
)

func GetDB() *sqlx.DB {
	onceDb.Do(func() {
		var err error
		cfg := config.GetConfig()
		dsn := "postgres://" + cfg.UserDB + ":" +
			cfg.PasswordDB + "@" + cfg.HostDB + ":" + cfg.PortDB + "/" + cfg.NameDB + "?sslmode=disable"
		db, err = sqlx.Open("pgx", dsn)
		if err != nil {
			log.Fatalln(err)
		}
		if err = db.Ping(); err != nil {
			log.Fatalln(err)
		}
		db.SetMaxOpenConns(maxConns)
	})
	return db
}
