package db

import (
	"context"
	"github.com/1hollow3/twittor-jupa/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func NewRegister(u models.User) (string, bool, error) {
	ctx, cancel:= context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("usernames")

	u.Password, _ = EncryptPassword(u.Password)

	result, err := col.InsertOne(ctx, u)

	if err != nil {
		return "", false, err
	}

	ObjID := result.InsertedID.(primitive.ObjectID).String()
	return ObjID, true, nil
}


