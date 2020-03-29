package scrapers

import (
	"opinion-reviews-scraper/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type DecathlonProductListResponse struct {
	ProductId int    `json:"product_id"`
	max_prize string `json:"max_price"`
	control   int    `json:"control"`
}

type DecathlonRecursiveAddScraperProduct struct {
	Config models.ScrapingConfig
	Index  models.ScrapingIndex
}

var maxPagesDecathlon = 300
var productBaseUrl = "https://www.decathlon.es/es/p/zapatillas-de-running-para-hombre-run-active-negras-y-amarillas/_/R-p-145826?mc="

func getProductList(baseUrl string) []DecathlonProductListResponse {
	response := []DecathlonProductListResponse{}

	res, err := http.Get(baseUrl)

	if err != nil {
		log.Fatal(err)
	}

	json.NewDecoder(res.Body).Decode(&response)
	return response
}

func composeProductUrl(productId string) string {
	//https://www.decathlon.es/es/p/zapatillas-de-running-para-hombre-run-active-negras-y-amarillas/_/R-p-145826?mc=8488621&c=ROJO
	return productBaseUrl + productId
}

func partialSave(productsReviews *[]models.ReviewScraped, config *models.ScrapingConfig, scrapingIndex *models.ScrapingIndex) {
	log.Println("Saving partial results")
	models.SaveMany(productsReviews, config)
	scrapingIndex.Save(*config)
}

func (scraper DecathlonRecursiveAddScraperProduct) ScrapReviewsInItems(baseUrl string, scrapingIndex *models.ScrapingIndex) {
	reviewsScraper := DecathlonReviewsScraperInProduct{scraper.Config}

	productList := getProductList(baseUrl)

	//results := []models.ReviewScraped{}

	for index, product := range productList {
		log.Println("-----------------------------")
		log.Printf("Starting scraping in page %d", index)
		log.Println("new page for decathlon")
		log.Println("-----------------------------")

		productId := strconv.Itoa(product.ProductId)
		url := composeProductUrl(productId)

		log.Println("Scraping url for reviews")
		log.Println(url)

		urlNew := UrlNew{url: url}
		productReviews := reviewsScraper.ScrapPage(urlNew)
		scrapingIndex.PageIndex = index
		partialSave(&productReviews, &scraper.Config, scrapingIndex)
	}

}
