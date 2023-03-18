package init

import (
	"log"
	"os"
	"sort"

	"github.com/Shalqarov/weather-bot/config"
	"github.com/Shalqarov/weather-bot/internal/bot"
	"github.com/Shalqarov/weather-bot/internal/database/postgres"
)

const (
	migrationPath = "./internal/database/postgres/migrations/"
)

func Config() {
	config.GetConfig()
}

func DB() {
	postgres.GetDB()
}

func Bot() {
	bot.GetBot()
}

func Migration() {
	f, err := os.Open(migrationPath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	files, err := f.Readdirnames(0)
	if err != nil {
		log.Fatal(err)
	}
	sort.Slice(files, func(i, j int) bool {
		return files[i] < files[j]
	})

	db := postgres.GetDB()
	for _, file := range files {
		path := migrationPath + file
		query, err := os.ReadFile(path)
		if err != nil {
			log.Fatalln(err)
		}

		_, err = db.Exec(string(query))
		if err != nil {
			log.Fatalln(err)
		}
	}
	log.Println("Migration succeded")
}
