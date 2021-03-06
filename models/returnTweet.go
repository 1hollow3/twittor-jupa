package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ReturnTweet struct {
	ID      primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID  string             `bson:"userID" json:"userID,omitempty"`
	Message string             `bson:"message" json:"message,omitempty"`
	Date    time.Time          `bson:"date" json:"date,omitempty"`
}
