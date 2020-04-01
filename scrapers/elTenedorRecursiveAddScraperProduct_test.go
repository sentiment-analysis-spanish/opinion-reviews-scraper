package scrapers

import (
	"opinion-reviews-scraper/models"
	"fmt"
	"testing"

	"github.com/joho/godotenv"
)

func TestRecursiveScraperElTenedor(t *testing.T) {
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}
	config := models.ScrapingConfig{UrlBase: "http://localhost:8000", ScraperId: "testScraperElpais", DeviceID: "testDeviceElpais"}
	index := models.ScrapingIndex{ScraperID: "test", PageIndex: 25}
	scraper := ElTenedorRecursiveAddScraperProduct{Config: config}

	//baseUrl := "https://www.amazon.es/gp/bestsellers/?ref_=nav_cs_bestsellers"
	baseUrl := "https://www.eltenedor.es/search/?cityId=328022"

	scraper.ScrapReviewsInItems(baseUrl, &index)

}
