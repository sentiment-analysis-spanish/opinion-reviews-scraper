package scrapers

import (
	"opinion-reviews-scraper/models"
	"fmt"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestReviewScraperEltenedor(t *testing.T) {
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}
	config := models.ScrapingConfig{UrlBase: "http://localhost:8000", ScraperId: "testScraperElpais", DeviceID: "testDeviceElpais"}
	scraper := ElTenedorReviewsScraper{Config: config}
	date := time.Date(2018, 05, 04, 0, 0, 0, 0, time.UTC)

	urlNew := UrlNew{url: "https://www.eltenedor.es/restaurante/mercado-de-la-reina-r6882", date: date}

	result := scraper.ScrapPage(urlNew)
	fmt.Println(result)

	//https://www.decathlon.es/es/ajax/asyncCartridgeLoad?contentPath=/content/Shared/Product%20Details%20Content/ReviewsFloor&params=A%3Dm-8400335%26mc%3D8400335

	assert.NotEmpty(t, result[0].Content, "Should fill tags")

}