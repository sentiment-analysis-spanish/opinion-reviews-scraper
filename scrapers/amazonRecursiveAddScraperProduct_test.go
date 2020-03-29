package scrapers

import (
	"fmt"
	"opinion-reviews-scraper/models"
	"testing"

	"github.com/joho/godotenv"
)

func TestRecursiveScraperAmazon(t *testing.T) {
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}
	config := models.ScrapingConfig{UrlBase: "http://localhost:8000", ScraperId: "testScraperElpais", DeviceID: "testDeviceElpais"}
	scraper := AmazonRecursiveAddScraperProduct{Config: config}

	//baseUrl := "https://www.amazon.es/gp/bestsellers/?ref_=nav_cs_bestsellers"
	baseUrl := "https://www.amazon.es/s?k=nintendo+switch&i=videogames&rh=n%3A599382031&dc&page=2&qid=1577956845&rnid=1703620031&ref=sr_pg_2"

	scraper.Start(baseUrl)

}
