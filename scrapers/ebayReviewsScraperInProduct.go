package scrapers

import (
	"opinion-reviews-scraper/models"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
)

type EbayReviewsScraperInProduct struct {
	Config models.ScrapingConfig
}

func RateTextToFloatEbay(rateText string) float64 {
	cleanRate := strings.ReplaceAll(rateText, "estrellas", "")
	splitted := strings.Split(cleanRate, " de ")
	rate, _ := strconv.ParseFloat(splitted[0], 64)
	rate = rate / 5
	return rate
}
func (scraper *EbayReviewsScraperInProduct) ScrapPage(urlNew UrlNew) []models.ReviewScraped {
	results := []models.ReviewScraped{}

	// Instantiate default collector
	c := colly.NewCollector(
	// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
	//colly.AllowedDomains("https://elpais.com/"),
	)

	c.Limit(&colly.LimitRule{
		Parallelism: 2,
		RandomDelay: 1 * time.Second,
	})

	c.OnHTML("div[itemprop]", func(e *colly.HTMLElement) {
		if e.Attr("itemprop") == "review" {
			result := models.ReviewScraped{}
			text := ""
			rateText := ""
			username := ""
			date := ""

			e.ForEach("p", func(_ int, elem *colly.HTMLElement) {
				if !strings.Contains(elem.Text, "Compra verificada:") {
					text = text + elem.Text + "\n"

				}
			})

			e.ForEach("div[role]", func(_ int, elem *colly.HTMLElement) {
				if elem.Attr("role") == "img" {
					rateText = elem.Attr("aria-label")
				}
			})

			e.ForEach("a[itemprop]", func(_ int, elem *colly.HTMLElement) {
				if elem.Attr("itemprop") == "author" {
					username = elem.Text
				}
			})

			e.ForEach("span[itemprop]", func(_ int, elem *colly.HTMLElement) {
				if elem.Attr("itemprop") == "datePublished" {
					date = elem.Text
				}
			})

			result.Content = text
			result.RateText = rateText
			result.Date = date
			result.Rate = RateTextToFloatEbay(rateText)
			result.User = username
			result.Source = "ebay"
			results = append(results, result)

			log.Println("obtained new review by user " + username + " rate: " + rateText)

		}
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

	c.Visit(urlNew.url)
	c.Wait()

	return results

}
