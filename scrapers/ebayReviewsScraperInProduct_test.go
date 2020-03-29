package scrapers

import (
	"github.com/stretchr/testify/assert"
	"opinion-reviews-scraper/models"
	"fmt"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

func TestReviewScraperEbay(t *testing.T) {
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}
	config := models.ScrapingConfig{UrlBase: "http://localhost:8000", ScraperId: "testScraperElpais", DeviceID: "testDeviceElpais"}
	scraper := EbayReviewsScraperInProduct{Config: config}
	date := time.Date(2018, 05, 04, 0, 0, 0, 0, time.UTC)

	urlNew := UrlNew{url: "https://www.ebay.es/itm/Cable-Original-100-Apple-Carga-Lightning-MD819-2m-Caja-Retail-iPhone-6s-7-X/254140893447?hash=item3b2bfa4907%3Ag%3AplIAAOSwK3ZbhCqY&_trkparms=%2526rpp_cid%253D5d00c682fe6ec94f90cc4945", date: date}

	result := scraper.ScrapPage(urlNew)
	fmt.Println(result)

	assert.NotEqual(t, result[0].RateText, "",  "Should fill tags")
	assert.NotEqual(t, result[0].Content, "",  "Should fill tags")

}
