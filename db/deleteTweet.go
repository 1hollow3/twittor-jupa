package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func DeleteTweet(ID string, UserID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("tweets")

	objID, _ := primitive.ObjectIDFromHex(ID)

	query := bson.M{
		"_id":    objID,
		"userID": UserID,
	}

	_, err := col.DeleteOne(ctx, query)
	return err
}
