package scrapers

import (
	"opinion-reviews-scraper/models"
	"fmt"
	"testing"

	"github.com/joho/godotenv"
)

func TestRecursiveScraperDecathlon(t *testing.T) {
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}
	config := models.ScrapingConfig{UrlBase: "http://localhost:8000", ScraperId: "testScraperElpais", DeviceID: "testDeviceElpais"}
	index := models.ScrapingIndex{ScraperID: "test"}
	scraper := DecathlonRecursiveAddScraperProduct{Config: config}

	//baseUrl := "https://www.amazon.es/gp/bestsellers/?ref_=nav_cs_bestsellers"
	baseUrl := "https://www.decathlon.es/static/override/bajada-permanente/productList.json?_=1578044145974"

	scraper.ScrapReviewsInItems(baseUrl, &index)

}
