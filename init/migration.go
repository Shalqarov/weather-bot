package init

import (
	"log"
	"os"
	"sort"

	"github.com/Shalqarov/weather-bot/internal/database/postgres"
)

func Migration(dir string) {
	f, err := os.Open(dir)
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
		path := dir + file
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
