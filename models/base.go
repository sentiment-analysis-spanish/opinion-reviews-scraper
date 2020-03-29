package models

import (
	"context"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

func init() {

	e := godotenv.Load()
	if e != nil {
		e := godotenv.Load("../.env")
		fmt.Print(e)
	}

	dbAddress := os.Getenv("database_url")
	dbName := os.Getenv("user_database_name")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI(dbAddress))
	db = client.Database(dbName)
}

func GetDB() *mongo.Database {
	return db
}
