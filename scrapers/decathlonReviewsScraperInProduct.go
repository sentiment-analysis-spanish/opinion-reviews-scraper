package scrapers

import (
	"opinion-reviews-scraper/models"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
)

type DecathlonReviewsScraperInProduct struct {
	Config models.ScrapingConfig
}

func RateTextToFloatDecathlon(rateText string) float64 {
	splitted := strings.Split(rateText, "/")
	rate, _ := strconv.ParseFloat(splitted[0], 64)
	rate = rate / 5
	return rate
}

var decathlonResponseIdentifier = "Atentamente para cualquier otra consulta"

func getAjaxUrl(url string) string {
	//https://www.decathlon.es/es/p/bicicleta-de-montana-rockrider-st-100-27-5-gris/_/R-p-192872?mc=8400335&c=GRIS

	//https://www.decathlon.es/es/ajax/asyncCartridgeLoad?contentPath=/content/Shared/Product%20Details%20Content/ReviewsFloor&params=A%3Dm-8400335%26mc%3D8400335
	//https://www.decathlon.es/es/ajax/asyncCartridgeLoad?contentPath=/content/Shared/Product%20Details/Supplemental%20Content/Floors%20Content/ReviewsFloor&params=A%3Dm-8400335%26mc%3D8400335
	prodId := strings.Split(url, "?mc=")[1]
	if strings.Contains(prodId, "&") {
		prodId = strings.Split(prodId, "&")[0]
	}

	return "https://www.decathlon.es/es/ajax/asyncCartridgeLoad?contentPath=/content/Shared/Product%20Details/Supplemental%20Content/Floors%20Content/ReviewsFloor&params=A%3Dm-" + prodId + "%26mc%3D" + prodId
}

func (scraper *DecathlonReviewsScraperInProduct) ScrapPage(urlNew UrlNew) []models.ReviewScraped {
	results := []models.ReviewScraped{}

	ajaxUrl := getAjaxUrl(urlNew.url)
	// Instantiate default collector
	c := colly.NewCollector(
	// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
	//colly.AllowedDomains("https://elpais.com/"),
	)

	c.Limit(&colly.LimitRule{
		Parallelism: 2,
		RandomDelay: 1 * time.Second,
	})

	c.OnHTML("article", func(e *colly.HTMLElement) {

			result := models.ReviewScraped{}
			text := ""
			rateText := ""
			username := ""
			date := ""
			title := ""

			e.ForEach("p[class]", func(_ int, elem *colly.HTMLElement) {
				if elem.Attr("class") == "review__text" && !strings.Contains(elem.Text, decathlonResponseIdentifier) {
					text = elem.Text

				}
			})

			e.ForEach("strong[class]", func(_ int, elem *colly.HTMLElement) {
				if elem.Attr("class") == "review__rating" {
					rateText = elem.Text
				}
			})

			e.ForEach("span[class]", func(_ int, elem *colly.HTMLElement) {
				if elem.Attr("class") == "review__date" {
					date = elem.Text
				}
			})

			e.ForEach("h4[class]", func(_ int, elem *colly.HTMLElement) {
				if elem.Attr("class") == "review__author-name" {
					username = elem.Text
				}
			})

			e.ForEach("h3[class]", func(_ int, elem *colly.HTMLElement) {
				if elem.Attr("class") == "title--main title--gtm-hero review__title" {
					title = elem.Text
				}
			})

			result.Content = title + text
			result.Title = title
			result.RateText = rateText
			result.Date = date
			result.Rate = RateTextToFloatDecathlon(rateText)
			result.User = username
			result.Source = "decathlon"
			results = append(results, result)
			log.Println("obtained new review by user " + username + " rate: " + rateText)


	})

	c.OnHTML("a", func(e *colly.HTMLElement) {
		url := e.Attr("href")
		if strings.Contains(url, "/itm/") {
			//fmt.Println("Recursively visiting")
			//c.Visit(url)
		}
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		log.Info("Visiting\n", r.URL.String())
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Info("Something went wrong:", err)
	})

	c.Visit(ajaxUrl)
	c.Wait()

	return results

}
