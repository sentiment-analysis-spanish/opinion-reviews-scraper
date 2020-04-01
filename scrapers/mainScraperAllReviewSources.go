package scrapers

import (
	"opinion-reviews-scraper/models"
	"opinion-reviews-scraper/utils"
	"sync"

	log "github.com/sirupsen/logrus"
)

type MainScraperAllReviewSources struct {
}

func (mainScraper MainScraperAllReviewSources) StartScraping(config models.ScrapingConfig) {

	scraperEbay := EbayRecursiveAddScraperProduct{Config: config}
	scraperDecathlon := DecathlonRecursiveAddScraperProduct{Config: config}
	scrapertripadvisor := TripAdvisorRecursiveAddScraperProduct{Config: config}
	scrapertripadvisorRestaurants := TripAdvisorRecursiveAddScraperProduct{Config: config}
	scraperEltenedor := ElTenedorRecursiveAddScraperProduct{Config: config}

	scrapAll := utils.StringInSlice("all", config.Scrapers)

	log.Info("using historicScrapers:")
	for {
		var wg sync.WaitGroup

		if utils.StringInSlice("eltenedor", config.Scrapers) || scrapAll {
			go mainScraper.ScrapOneIteration(scraperEltenedor, "eltenedor", config, &wg)
			wg.Add(1)
		}
		if utils.StringInSlice("tripadvisor-restaurant", config.Scrapers) || scrapAll {
			go mainScraper.ScrapOneIteration(scrapertripadvisorRestaurants, "tripadvisor-restaurant", config, &wg)
			wg.Add(1)
		}

		if utils.StringInSlice("ebay", config.Scrapers) || scrapAll {
			go mainScraper.ScrapOneIteration(scraperEbay, "ebay", config, &wg)
			wg.Add(1)
		}

		if utils.StringInSlice("decathlon", config.Scrapers) || scrapAll {
			go mainScraper.ScrapOneIteration(scraperDecathlon, "decathlon", config, &wg)
			wg.Add(1)
		}

		if utils.StringInSlice("tripadvisor", config.Scrapers) || scrapAll {
			go mainScraper.ScrapOneIteration(scrapertripadvisor, "tripadvisor", config, &wg)
			wg.Add(1)
		}

		wg.Wait()
		log.Info("-------------------------------------------------------------------------------------------------")
		log.Info("-------------------Finished one iteration, all news from page scraped----------------------------")
		log.Info("-------------------------------------------------------------------------------------------------")
	}

}

func (mainScraper *MainScraperAllReviewSources) ScrapOneIteration(scraper RecursiveReviewsScraper, source string, config models.ScrapingConfig, wg *sync.WaitGroup) {
	defer wg.Done()
	log.Info("starting scraping using " + source)
	scrapingIndex := models.ScrapingIndex{}
	scrapingIndex.GetCurrentIndex(config, source)
	scrapingIndex.UpdateUrls(config, source)

	index := scrapingIndex.UrlIndex
	if index >= len(scrapingIndex.StartingUrls) {
		index = 0
	}

	log.Printf("starting with url number %d", index)

	nextUrl := scrapingIndex.StartingUrls[index]
	scraper.ScrapReviewsInItems(nextUrl, &scrapingIndex)

	//models.SaveMany(results, config)

	index = index + 1

	scrapingIndex.UrlIndex = index

	scrapingIndex.Save(config)
}
