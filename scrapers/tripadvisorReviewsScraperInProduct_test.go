package scrapers

import (
	"opinion-reviews-scraper/models"
	"fmt"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestReviewScraperTripadvisor(t *testing.T) {
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}
	config := models.ScrapingConfig{UrlBase: "http://localhost:8000", ScraperId: "testScraperElpais", DeviceID: "testDeviceElpais"}
	scraper := TripAdvisorReviewsScraperInProduct{Config: config}
	date := time.Date(2018, 05, 04, 0, 0, 0, 0, time.UTC)

	urlNew := UrlNew{url: "https://www.tripadvisor.es/Hotel_Review-g311415-d597947-Reviews-or1455-The_St_Regis_Bora_Bora_Resort-Bora_Bora_Society_Islands.html#REVIEWS", date: date}

	result := scraper.ScrapPage(urlNew)
	fmt.Println(result)

	//https://www.decathlon.es/es/ajax/asyncCartridgeLoad?contentPath=/content/Shared/Product%20Details%20Content/ReviewsFloor&params=A%3Dm-8400335%26mc%3D8400335

	assert.NotEmpty(t, result[0].Content, "Should fill tags")

}


//https://www.tripadvisor.es/Restaurant_Review-g187514-d19740783-Reviews-Nebak_Jatetxea_Madrid-Madrid.html
func TestReviewScraperTripadvisorRestaurants(t *testing.T) {
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}
	config := models.ScrapingConfig{UrlBase: "http://localhost:8000", ScraperId: "testScraperElpais", DeviceID: "testDeviceElpais"}
	scraper := TripAdvisorReviewsScraperInProduct{Config: config}
	date := time.Date(2018, 05, 04, 0, 0, 0, 0, time.UTC)

	urlNew := UrlNew{url: "https://www.tripadvisor.es/Restaurant_Review-g187514-d19740783-Reviews-Nebak_Jatetxea_Madrid-Madrid.html", date: date}

	result := scraper.ScrapPage(urlNew)
	fmt.Println(result)

	assert.NotEmpty(t, result[0].Content, "Should fill tags")

}