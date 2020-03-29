package main

import (
	"fmt"
	"io"
	"opinion-reviews-scraper/models"
	"opinion-reviews-scraper/scrapers"
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

const logPath = "logs.txt"

func main() {
	godotenv.Load()

	InitLogger()

	config := &models.ScrapingConfig{}
	config.CreateFromJson()
	fmt.Println(config)
	if config == nil {
		panic("No config")
	}

	var mainScraper scrapers.MainScraper

	mainScraper = scrapers.MainScraperAllReviewSources{}

	mainScraper.StartScraping(*config)
}

func InitLogger() {

	file, err := os.OpenFile("logs.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalln("Failed to open log file")
	}
	mw := io.MultiWriter(os.Stdout, file)
	log.SetOutput(mw)

}
