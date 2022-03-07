package db

import (
	"context"
	"time"

	"github.com/1hollow3/twittor-jupa/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func  ModifyUser(user models.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("usernames")

	register := make(map[string]interface{})
	if len(user.Name) > 0 {
		register["name"] = user.Name
	}
	if len(user.Lastname) > 0 {
		register["lastname"] = user.Lastname
	}
	register["birthDay"] = user.BirthDay
	if len(user.Avatar) > 0 {
		register["avatar"] = user.Avatar
	}
	if len(user.Banner) > 0 {
		register["banner"] = user.Banner
	}
	if len(user.Bio) > 0 {
		register["bio"] = user.Bio
	}
	if len(user.Location) > 0 {
		register["location"] = user.Location
	}
	if len(user.WebSite) > 0 {
		register["website"] = user.WebSite
	}

	obID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": obID}
	updateString := bson.M{
		"$set": register,
	}

	_, err := col.UpdateOne(ctx, filter, updateString)
	if err != nil {
		return false, err
	}

	return true, nil
}
