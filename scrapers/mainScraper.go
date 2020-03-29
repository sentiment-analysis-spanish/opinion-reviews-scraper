package scrapers

import "opinion-reviews-scraper/models"

type MainScraper interface {
	StartScraping(config models.ScrapingConfig)
}
