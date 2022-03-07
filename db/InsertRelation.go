package db

import (
	"context"
	"fmt"
	"github.com/1hollow3/twittor-jupa/models"
	"time"
)

func InsertRelation(t models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("relation")

	var prev models.Relation
	col.FindOne(ctx, t).Decode(&prev)

	if len(prev.UserID)> 0 {
		return false, fmt.Errorf("the relation exist already ")
	}

	_, err := col.InsertOne(ctx, t)
	if err != nil {
		return false, err
	}

	return true, nil

}
