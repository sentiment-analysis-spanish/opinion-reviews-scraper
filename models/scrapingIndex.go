package models

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type ScrapingIndex struct {
	DateScraping  time.Time `json:"date_scraping" bson:"date_scraping"`
	UrlIndex      int       `json:"url_index" bson:"url_index"`
	PageIndex     int       `json:"page_index" bson:"page_index"`
	ReviewsSource string    `json:"reviews_source" bson:"reviews_source"`
	StartingUrls  []string  `json:"startingUrls" bson:"startingUrls"`
	ScraperID     string    `json:"scraper_id" bson:"scraper_id"`
	DeviceID      string    `json:"device_id" bson:"device_id"`
}

var collectionNameScrapingIndex = "ScrapingIndex"

func (scrapingIndex *ScrapingIndex) Save(config ScrapingConfig) error {
	log.Println("Updating scraping index")
	db := GetDB()
	scrapingIndex.DateScraping = time.Now()
	collection := db.Collection(collectionNameScrapingIndex)
	filter := bson.M{"review_source": scrapingIndex.ReviewsSource}
	update := bson.M{"$set": scrapingIndex}
	_, err := collection.UpdateOne(context.Background(), filter, update, options.Update().SetUpsert(true))

	if err != nil {
		log.Fatal(err)
		return err

	}
	return nil
}

func (scrapingIndex *ScrapingIndex) Delete() error {
	db := GetDB()
	collection := db.Collection(collectionNameScrapingIndex)
	filter := bson.M{"scraper_id": scrapingIndex.ScraperID, "reviews_source": scrapingIndex.ReviewsSource}
	_, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
		return err

	}
	return nil
}

func (scrapingIndex *ScrapingIndex) GetCurrentIndex(config ScrapingConfig, source string) error {
	err := scrapingIndex.getCurrentIndexFromService(config, source)
	if err != nil {
		log.Println(err)
	}
	if scrapingIndex.ScraperID == "" {
		log.Println("since the index is empty we fill it with config values")
		scrapingIndex.ScraperID = config.ScraperId
		scrapingIndex.DeviceID = config.DeviceID
		scrapingIndex.DateScraping = time.Now()
		scrapingIndex.UrlIndex = 0
		scrapingIndex.PageIndex = 1
		scrapingIndex.ReviewsSource = source
		scrapingIndex.StartingUrls = config.InitialUrls[source]
		scrapingIndex.Save(config)
	}
	return nil
}

func (scrapingIndex *ScrapingIndex) UpdateUrls(config ScrapingConfig, source string) error {
	err := scrapingIndex.getCurrentIndexFromService(config, source)
	if err != nil {
		log.Println(err)
	}
	log.Println("Updating initial urls for scraping in " + source)
	scrapingIndex.StartingUrls = config.InitialUrls[source]
	scrapingIndex.Save(config)

	return nil
}

func (scrapingIndex *ScrapingIndex) getCurrentIndexFromService(config ScrapingConfig, reviewSource string) error {
	log.Println("Retrieving scraping index")
	db := GetDB()
	collection := db.Collection(collectionNameScrapingIndex)
	err := collection.FindOne(context.Background(), bson.M{"review_source": reviewSource}).Decode(scrapingIndex)
	if err != nil { //User not found!
		return err
	}
	return nil
}
