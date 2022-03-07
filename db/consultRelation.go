package db

import (
	"context"
	"github.com/1hollow3/twittor-jupa/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func ConsultRelation(t models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("relation")

	query := bson.M{
		"userID":        t.UserID,
		"relatedUserID": t.RelatedUserID,
	}
	var result models.Relation
	err := col.FindOne(ctx, query).Decode(&result)
	if err != nil {
		return false, err
	}

	return true, nil

}
