package scrapers

import (
	"fmt"
	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
	"opinion-reviews-scraper/models"
	"strings"
	"time"
)

type ElTenedorRecursiveAddScraperProduct struct {
	Config models.ScrapingConfig
	Index  models.ScrapingIndex
}

func (scraper ElTenedorRecursiveAddScraperProduct) ScrapReviewsInItems(baseUrl string, scrapingIndex *models.ScrapingIndex) {
	//results := []models.ReviewScraped{}
	urlsPending := []UrlNew{}
	reviewsScraper := ElTenedorReviewsScraper{scraper.Config}

	// Instantiate default collector
	c := colly.NewCollector(
	// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
	//colly.AllowedDomains("https://elpais.com/"),
	)
	c.OnHTML("a[class]", func(e *colly.HTMLElement) {
			url := e.Attr("href")
			if strings.Contains(url, "/restaurante/") {
				date := time.Now()
				urlScrap := UrlNew{url: "https://www.eltenedor.es" + url, date: date}
				fmt.Println(url)
				urlsPending = append(urlsPending, urlScrap)
			}
	})

	c.OnHTML("ul", func(e *colly.HTMLElement) {
		if e.Attr("class")=="_1fOA6" {
			e.ForEach("li", func(_ int, elem *colly.HTMLElement) {
				elem.ForEach("a", func(_ int, elem2 *colly.HTMLElement) {
					url := "https://www.eltenedor.es" + elem2.Attr("href") //strings.Replace(e.Attr("href"), "/", "", 1)
					c.Visit(url)
				})
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

func (scraper ElTenedorRecursiveAddScraperProduct) scrapAllReviewsInUrl(urlbase UrlNew, reviewsScraper *ElTenedorReviewsScraper, out chan []models.ReviewScraped) []models.ReviewScraped {
	results := reviewsScraper.ScrapPage(urlbase)

	out <- results
	return results
}
