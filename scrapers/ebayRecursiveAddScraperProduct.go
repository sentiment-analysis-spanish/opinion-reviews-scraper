package scrapers

import (
	"opinion-reviews-scraper/models"
	"fmt"
	"strings"
	"time"

	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
)

type EbayRecursiveAddScraperProduct struct {
	Config models.ScrapingConfig
	Index  models.ScrapingIndex
}

var maxPages = 200

func getUrl(baseUrl string, page int) (url string) {
	baseUrl = strings.ReplaceAll(baseUrl, "?rt=nc", "")
	if strings.Contains(baseUrl, "&_pgn=") {
		url = strings.Split(baseUrl, "&_pgn=")[0] + "?_pgn=" + fmt.Sprintf("%d", page)
		return url
	} else {
		url = strings.Split(baseUrl, "?_pgn=")[0] + "?_pgn=" + fmt.Sprintf("%d", page)
		return url
	}

}

func (scraper EbayRecursiveAddScraperProduct) ScrapReviewsInItems(baseUrl string, scrapingIndex *models.ScrapingIndex) {
	//results := []models.ReviewScraped{}

	for scrapingIndex.PageIndex < maxPages {
		out1 := make(chan []models.ReviewScraped)
		out2 := make(chan []models.ReviewScraped)
		out3 := make(chan []models.ReviewScraped)

		page1 := scrapingIndex.PageIndex
		page2 := scrapingIndex.PageIndex + 1
		page3 := scrapingIndex.PageIndex + 2

		go scraper.scrapAllReviewsInUrl(baseUrl, page1, out1)
		go scraper.scrapAllReviewsInUrl(baseUrl, page2, out2)
		go scraper.scrapAllReviewsInUrl(baseUrl, page3, out3)

		resultsInPage1, resultsInPage2, resultsInPage3 := <-out1, <-out2, <-out3
		models.SaveMany(&resultsInPage1, &scraper.Config)
		models.SaveMany(&resultsInPage2, &scraper.Config)
		models.SaveMany(&resultsInPage3, &scraper.Config)

		//results = append(results, resultsInPage1...)
		//results = append(results, resultsInPage2...)

		scraper.updateIndex(scrapingIndex.PageIndex+3, scrapingIndex)
	}

	scraper.updateIndex(1, scrapingIndex)

	//return results

}

func (scraper EbayRecursiveAddScraperProduct) scrapAllReviewsInUrl(urlbase string, page int, out chan []models.ReviewScraped) []models.ReviewScraped {

	reviewsScraper := EbayReviewsScraperInProduct{scraper.Config}
	urlvisit := getUrl(urlbase, page)

	results := []models.ReviewScraped{}
	// Instantiate default collector
	c := colly.NewCollector(
	// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
	//colly.AllowedDomains("https://elpais.com/"),
	)

	log.Println("-----------------------------")
	log.Printf("Starting scraping in page %d", page)
	log.Println("new page for ebay")
	log.Println("-----------------------------")

	c.Limit(&colly.LimitRule{
		Parallelism: 4,
		RandomDelay: 1 * time.Second,
	})

	c.OnHTML("a", func(e *colly.HTMLElement) {
		url := e.Attr("href")
		if strings.Contains(url, "/itm/") {
			date := time.Now()
			urlScrap := UrlNew{url: url, date: date}
			scrapedreviews := reviewsScraper.ScrapPage(urlScrap)
			results = append(results, scrapedreviews...)
		}
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		log.Info("Visiting\n", r.URL.String())
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Info("Something went wrong:", err)
	})

	c.Visit(urlvisit)
	out <- results
	return results

}

func (scraper EbayRecursiveAddScraperProduct) updateIndex(page int, scrapingIndex *models.ScrapingIndex) {
	scrapingIndex.PageIndex = page
	scrapingIndex.Save(scraper.Config)
}
