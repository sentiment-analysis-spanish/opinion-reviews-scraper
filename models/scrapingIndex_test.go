package models

import (
	"testing"
)

func TestGetScrapingIndex(t *testing.T) {

	config := ScrapingConfig{}
	config.CreateFromJson()

}

func TestPostScrapingIndex(t *testing.T) {
	config := ScrapingConfig{}
	config.CreateFromJson()

	index := ScrapingIndex{}
	index.NewsPaper = "elpais"
	index.DeviceID = config.DeviceID
	index.ScraperID = config.ScraperId

}
