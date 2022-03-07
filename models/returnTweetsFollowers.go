package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ReturnTweetsFollowers struct {
	ID            primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID        string             `bson:"userID" json:"userID,omitempty"`
	RelatedUserID string             `bson:"relatedUserID" json:"userRelatedID,omitempty"`
	Tweet         struct {
		Message string    `bson:"message" json:"message,omitempty"`
		Date    time.Time `bson:"date" json:"date,omitempty"`
		ID      string    `bson:"id" json:"id,omitempty"`
	}
}
