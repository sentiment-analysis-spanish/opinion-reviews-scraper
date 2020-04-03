package scrapers

import (
	"opinion-reviews-scraper/models"
	"fmt"
	"strings"
	"time"

	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
)

type FilmaffinityRecursiveAddScraperProduct struct {
	Config models.ScrapingConfig
	Index  models.ScrapingIndex
}

func (scraper FilmaffinityRecursiveAddScraperProduct) ScrapReviewsInItems(baseUrl string, scrapingIndex *models.ScrapingIndex) {
	//results := []models.ReviewScraped{}
	urlsPending := []UrlNew{}
	reviewsScraper := FilmaffinityReviewsScraper{scraper.Config}

	// Instantiate default collector
	c := colly.NewCollector(
	// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
	//colly.AllowedDomains("https://elpais.com/"),
	)

	c.Limit(&colly.LimitRule{
		Parallelism: 2,
		RandomDelay: 1 * time.Second,
	})

	c.OnHTML("a", func(e *colly.HTMLElement) {
		if strings.Contains(e.Attr("href"), "/es/film") {
			url := e.Attr("href")
			date := time.Now()
			urlScrap := UrlNew{url: url, date: date}
			fmt.Println(url)
			urlsPending = append(urlsPending, urlScrap)
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

	log.Println("Collected pages")
	log.Println(urlsPending)

	for scrapingIndex.PageIndex < len(urlsPending)-3 {
		log.Println("---------------")
		log.Printf("Scraping page %d", scrapingIndex.PageIndex)
		log.Println("---------------")

		out1 := make(chan []models.ReviewScraped)
		out2 := make(chan []models.ReviewScraped)
		out3 := make(chan []models.ReviewScraped)

		url1 := urlsPending[scrapingIndex.PageIndex]
		url2 := urlsPending[scrapingIndex.PageIndex+1]
		url3 := urlsPending[scrapingIndex.PageIndex+2]

		go scraper.scrapAllReviewsInUrl(url1, &reviewsScraper, out1)
		go scraper.scrapAllReviewsInUrl(url2, &reviewsScraper, out2)
		go scraper.scrapAllReviewsInUrl(url3, &reviewsScraper, out3)

		resultsInPage1, resultsInPage2, resultsInPage3 := <-out1, <-out2, <-out3
		models.SaveMany(&resultsInPage1, &scraper.Config)
		models.SaveMany(&resultsInPage2, &scraper.Config)
		models.SaveMany(&resultsInPage3, &scraper.Config)

		//results = append(results, resultsInPage1...)
		//results = append(results, resultsInPage2...)

		scrapingIndex.PageIndex = scrapingIndex.PageIndex + 3
		scrapingIndex.Save(scraper.Config)

	}
	scrapingIndex.PageIndex = 0
	scrapingIndex.Save(scraper.Config)

	//return results

}

func (scraper FilmaffinityRecursiveAddScraperProduct) scrapAllReviewsInUrl(urlbase UrlNew, reviewsScraper *FilmaffinityReviewsScraper, out chan []models.ReviewScraped) []models.ReviewScraped {
	results := reviewsScraper.ScrapPage(urlbase)

	out <- results
	return results
}
