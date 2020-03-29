package scrapers

import (
	"opinion-reviews-scraper/models"
	"fmt"
	"testing"

	"github.com/joho/godotenv"
)

func TestRecursiveScraperTripadvisor(t *testing.T) {
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}
	config := models.ScrapingConfig{UrlBase: "http://localhost:8000", ScraperId: "testScraperElpais", DeviceID: "testDeviceElpais"}
	index := models.ScrapingIndex{ScraperID: "test", PageIndex: 25}
	scraper := TripAdvisorRecursiveAddScraperProduct{Config: config}

	//baseUrl := "https://www.amazon.es/gp/bestsellers/?ref_=nav_cs_bestsellers"
	baseUrl := "https://www.tripadvisor.es/Hotels-g187514-Madrid-Hotels.html"

	scraper.ScrapReviewsInItems(baseUrl, &index)

}
