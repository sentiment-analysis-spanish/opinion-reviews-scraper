package scrapers

import (
	"opinion-reviews-scraper/models"
)

type RecursiveReviewsScraper interface {
	ScrapReviewsInItems(url string, scrapingIndex *models.ScrapingIndex)
}
