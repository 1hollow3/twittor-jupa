package db

import (
	"context"
	"github.com/1hollow3/twittor-jupa/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func UserExist(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("usernames")

	query := bson.M{"email": email}

	var result models.User

	err := col.FindOne(ctx, query).Decode(&result)
	ID := result.ID.Hex()
	if err != nil {
		return result, false, ID
	}
	return result, true, ID

}
