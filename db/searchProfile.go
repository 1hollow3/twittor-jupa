package db

import (
	"context"
	"fmt"
	"time"

	"github.com/1hollow3/twittor-jupa/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SearchProfile(ID string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("usernames")

	var profile models.User
	objID, _ := primitive.ObjectIDFromHex(ID)

	query := bson.M{"_id": objID}

	err := col.FindOne(ctx, query).Decode(&profile)
	profile.Password = ""
	if err != nil {
		fmt.Println("Register not found " + err.Error())
		return profile, err
	}
	return profile, nil
}
