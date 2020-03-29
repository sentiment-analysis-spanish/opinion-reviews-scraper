package models

import (
	"encoding/json"
	"io/ioutil"
	"os"

	log "github.com/sirupsen/logrus"
)

type ScrapingConfig struct {
	UrlBase     string              `json:"url_base" bson:"url_base"`
	DeviceID    string              `json:"device_id" bson:"device_id"`
	ScraperId   string              `json:"scraper_id" bson:"scraper_id"`
	AppID       string              `json:"app_id" bson:"app_id"`
	Scrapers    []string            `json:"scrapers" bson:"scrapers"`
	InitialUrls map[string][]string `json:"initial_urls" bson:"initial_urls"`
}

func (config *ScrapingConfig) CreateFromJson() {
	jsonFile, err := os.Open("scrapingConfig.json")
	if err != nil {
		log.Error(err)
	}
	log.Info("Successfully Opened users.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &config)
}
