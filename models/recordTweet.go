package models

import "time"

type RecordTweet struct {
	UserID  string    `bson:"userID" json:"user_id,omitempty"`
	Message string    `bson:"message" json:"message,omitempty"`
	Date    time.Time `bson:"date" json:"date,omitempty"`
}
