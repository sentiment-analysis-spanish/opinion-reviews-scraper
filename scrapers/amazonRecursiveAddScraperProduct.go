package scrapers

import (
	"opinion-reviews-scraper/models"
	"fmt"
	"strings"
	"time"

	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
)

type AmazonRecursiveAddScraperProduct struct {
	Config models.ScrapingConfig
}

func (scraper *AmazonRecursiveAddScraperProduct) Start(baseUrl string) {
	reviewsScraper := AmazonReviewsScraperInProduct{scraper.Config}

	// Instantiate default collector
	c := colly.NewCollector(
	// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
	//colly.AllowedDomains("https://elpais.com/"),
	)
	c.OnHTML("div[data-component-type]", func(e *colly.HTMLElement) {
		if e.Attr("data-component-type") == "sp-sponsored-result" {
			e.ForEach("a[href]", func(_ int, elem *colly.HTMLElement) {
				url := elem.Attr("href")

				date := time.Now()
				urlScrap := UrlNew{url: "https://amazon.es/" + url, date: date}
				fmt.Println(url)
				reviewsScraper.ScrapPage(urlScrap)
				//e.Request.Visit(url)
			})
		}
	})

	c.OnHTML("ul[class]", func(e *colly.HTMLElement) {
		if e.Attr("class") == "a-pagination" {
			e.ForEach("a[href]", func(_ int, elem *colly.HTMLElement) {
				if strings.Contains(elem.Text, "Siguiente") {
					url := "https://amazon.es/" + elem.Attr("href")
					fmt.Println("next page:")
					fmt.Println(url)
					//e.Request.Visit(url)
				}
			})
		}
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		log.Info("Visiting\n", r.URL.String())
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Info("Something went wrong:", err)
	})

	c.Visit(baseUrl)

}
