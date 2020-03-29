package models

import (
	"fmt"
	"testing"
)

func TestPostNew(t *testing.T) {
	config := ScrapingConfig{}
	config.CreateFromJson()

	newScraped := ReviewScraped{}
	newScraped.Url = "http://test"
	newScraped.Content = "test"
	newScraped.ScraperID = config.ScraperId
	err := newScraped.Save(config)
	if err != nil {
		fmt.Println(err)
	}
	//assert.NotEqual(t, nil, err, "no error")
	fmt.Println(newScraped)

}
