package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"
	"strconv"
	"time"

	"xkcd/pkg/config"
	"xkcd/pkg/database"
	"xkcd/pkg/words"
	"xkcd/pkg/xkcd"
)

func main() {
	oFlag := flag.Bool("o", false, "Output file")
	nFlag := flag.Int("n", 10, "A number to view comics")
	cFlag := flag.Bool("c", false, "Print config")
	aFlag := flag.Bool("a", false, "Download all comics")
	flag.Parse()

	cfg := config.NewConfig()
	filePath, err := filepath.Abs("./config.yaml")
	if err != nil {
		log.Fatalf("Error getting absolute path: %v", err)
	}
	if err := cfg.ParseYAML(filePath); err != nil {
		log.Fatalf("Error parsing YAML file: %v", err)
	}

	if *cFlag {
		fmt.Printf("%-24s => %s\n", "SERVER URL", cfg.Xkcd.Source)
		fmt.Printf("%-24s => %s\n", "DATABASE", cfg.Xkcd.DbFile)
		fmt.Printf("%-24s => %d\n", "DATABASE SIZE LIMIT", cfg.Xkcd.DbSize)
		return
	}

	if *oFlag {
		database.ReadFirstNComics(cfg.Xkcd.DbFile, *nFlag)
		return
	}

	if *aFlag {
		fmt.Println("Fetching all comics...")
	} else {
		fmt.Printf("Fetching first %d comics...\n", cfg.Xkcd.DbSize)
		db := make(database.ComicsDatabase)
		for i := 1; i < cfg.Xkcd.DbSize; i++ {
			fetchComics(i, cfg, db)
		}

		// Сохранение базы данных в файл
		if err := database.SaveToFile(cfg.Xkcd.DbFile, db); err != nil {
			fmt.Printf("Failed to save comics to database: %v", err)
			return
		}
	}

	if oFlag != nil {
		fmt.Printf("%v\n", oFlag)
	}

	if nFlag != nil {
		fmt.Printf("%v\n", nFlag)
	}

}

func fetchComics(id int, cfg *config.Config, db database.ComicsDatabase) {
	client := xkcd.NewXkcdClient(30 * time.Second)
	url := cfg.Xkcd.Source

	// Получение данных с сервера
	comics, err := client.GetComics(url, id)
	if err != nil {
		fmt.Printf("Failed to fetch data: %v", err)
		return
	}

	// Стемминг
	stammer := words.NewStammer()
	stammedWords, err := stammer.Stem(comics.Transcript)
	if err != nil {
		fmt.Printf("Failed to stamming data: %v", err)
		return
	}

	// Формирование объекта, который запишется в БД
	record := database.ComicInfo{
		URL:      comics.Img,
		Keywords: stammedWords,
	}
	db[strconv.Itoa(id)] = record
}
