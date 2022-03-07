package db

import (
	"context"
	"github.com/1hollow3/twittor-jupa/models"
	"time"
)

func DeleteRelation(t models.Relation) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("relation")

	_, err := col.DeleteOne(ctx, t)
	return err
}
