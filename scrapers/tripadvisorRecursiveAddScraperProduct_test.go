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

func TestRecursiveScraperTripadvisorRestaurantes(t *testing.T) {
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}
	config := models.ScrapingConfig{UrlBase: "http://localhost:8000", ScraperId: "testScraperElpais", DeviceID: "testDeviceElpais"}
	index := models.ScrapingIndex{ScraperID: "test", PageIndex: 25}
	scraper := TripAdvisorRecursiveAddScraperProduct{Config: config}

	//baseUrl := "https://www.amazon.es/gp/bestsellers/?ref_=nav_cs_bestsellers"
	//baseUrl := "https://www.tripadvisor.es/Restaurants-g187514-Madrid.html"

	baseUrl:="https://www.tripadvisor.es/RestaurantSearch?Action=PAGE&ajax=1&availSearchEnabled=true&sortOrder=popularity&geo=187514&itags=10591&eaterydate=2020_04_01&date=2030-04-02&time=20%3A00%3A00&people=2&o=a10"

	scraper.ScrapReviewsInItems(baseUrl, &index)

}