package scrapers

import (
	"opinion-reviews-scraper/models"
	"fmt"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

func TestReviewScraper(t *testing.T) {
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}
	config := models.ScrapingConfig{UrlBase: "http://localhost:8000", ScraperId: "testScraperElpais", DeviceID: "testDeviceElpais"}
	scraper := AmazonReviewsScraperInProduct{Config: config}
	date := time.Date(2018, 05, 04, 0, 0, 0, 0, time.UTC)

	urlNew := UrlNew{url: "https://www.amazon.es/dp/B074H77NCM/ref=cm_gf_aAN_i2_d_p0_c0_qd0____________________Ud7u9KMqgwmoycr7N1ve?th=1", date: date}

	result := scraper.ScrapPage(urlNew)
	fmt.Println(result)

	//assert.NotEmpty(t, result.Tags, "Should fill tags")

}
