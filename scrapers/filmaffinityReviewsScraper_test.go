package scrapers

import (
	"opinion-reviews-scraper/models"
	"fmt"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestReviewScraperFilmaffinity(t *testing.T) {
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}
	config := models.ScrapingConfig{UrlBase: "http://localhost:8000", ScraperId: "testScraperElpais", DeviceID: "testDeviceElpais"}
	scraper := FilmaffinityReviewsScraper{Config: config}
	date := time.Date(2018, 05, 04, 0, 0, 0, 0, time.UTC)

	urlNew := UrlNew{url: "https://www.filmaffinity.com/es/film825296.html", date: date}

	result := scraper.ScrapPage(urlNew)
	fmt.Println(result[0].User)

	//https://www.decathlon.es/es/ajax/asyncCartridgeLoad?contentPath=/content/Shared/Product%20Details%20Content/ReviewsFloor&params=A%3Dm-8400335%26mc%3D8400335

	assert.NotEmpty(t, result[0].Content, "Should fill tags")

}