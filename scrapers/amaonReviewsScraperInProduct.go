package scrapers

import (
	"opinion-reviews-scraper/models"
	"fmt"
	"time"

	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
)

type UrlNew struct {
	url  string
	date time.Time
}
type AmazonReviewsScraperInProduct struct {
	Config models.ScrapingConfig
}

func (scraper *AmazonReviewsScraperInProduct) ScrapPage(urlNew UrlNew) models.ReviewScraped {
	result := models.ReviewScraped{}
	result.Tags = []string{}

	// Instantiate default collector
	c := colly.NewCollector(
	// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
	//colly.AllowedDomains("https://elpais.com/"),
	)
	c.OnHTML("div", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)

	})

	c.OnHTML("div[data-hook]", func(e *colly.HTMLElement) {
		if e.Attr("data-hook") == "review" {
			e.ForEach("span[data-hook]", func(_ int, elem *colly.HTMLElement) {
				if elem.Attr("data-hook") == "review-body" {
					fmt.Println(elem.Text)
				}
			})

			e.ForEach("span[class]", func(_ int, elem *colly.HTMLElement) {
				if elem.Attr("class") == "a-icon-alt" {
					fmt.Println("------")
					fmt.Println(elem.Text)
				}
			})
			//fmt.Println(e.Text)
		}
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		log.Info("Visiting", r.URL.String())
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Info("Something went wrong:", err)
	})

	c.Visit(urlNew.url)

	return result

}
