package models

import (
	"fmt"
	"testing"
)

func TestScrapingConfig(t *testing.T) {

	config := ScrapingConfig{}
	config.CreateFromJson()
	fmt.Println(config)
	//assert.NotEqual(t, nil, config.UrlBase, "OK response is expected")

}
