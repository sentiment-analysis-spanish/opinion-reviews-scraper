package models

import (
	"context"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type ReviewScraped struct {
	RateText  string   `json:"reate_text" bson:"reate_text"`
	Source    string   `json:"source" bson:"source"`
	User      string   `json:"user" bson:"user"`
	Rate      float64  `json:"rate" bson:"rate"`
	Date      string   `json:"date" bson:"date"`
	Content   string   `json:"content" bson:"content"`
	Title     string   `json:"title" bson:"title"`
	Tags      []string `json:"tags" bson:"tags"`
	Url       string   `json:"url" bson:"url"`
	ScraperID string   `json:"scraper_id" bson:"scraper_id"`
	ID        string   `json:"id" bson:"id"`
}

func (newScraped *ReviewScraped) Save(config *ScrapingConfig) error {
	db := GetDB()
	collection := db.Collection("Reviews")
	filter := bson.M{"content": newScraped.Content, "user": newScraped.User}
	update := bson.M{"$set": newScraped}
	_, err := collection.UpdateOne(context.Background(), filter, update, options.Update().SetUpsert(true))

	if err != nil {
		log.Fatal(err)
		return err

	}
	return nil
}

func SaveMany(newScraped *[]ReviewScraped, config *ScrapingConfig) error {
	for _, scraped := range *newScraped {
		scraped.Save(config)
	}
	return nil
	//err = json.NewDecoder(rs.Body).Decode(&scrapingIndex)
}
