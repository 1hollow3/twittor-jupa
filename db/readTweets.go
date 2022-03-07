package db

import (
	"context"
	"github.com/1hollow3/twittor-jupa/models"
	"go.mongodb.org/mongo-driver/bson"
	options2 "go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func ReadTweets(ID string, page int64) ([]*models.ReturnTweet, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("tweets")

	var result []*models.ReturnTweet

	query := bson.M{"userID": ID}
	options := options2.Find()
	options.SetSort(bson.D{{Key: "date", Value: -1}})
	options.SetSkip((page - 1) * 20)
	options.SetLimit(20)

	pointer, err := col.Find(ctx, query, options)
	if err != nil {
		log.Fatal(err.Error())
		return result, false
	}

	for pointer.Next(context.TODO()) {

		var tweet models.ReturnTweet
		err := pointer.Decode(&tweet)
		if err != nil {
			return result, false
		}
		result = append(result, &tweet)
	}

	return result, true
}
