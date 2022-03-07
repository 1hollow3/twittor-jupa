package models

type Relation struct{
	UserID string `bson:"userID" json:"userID"`
	RelatedUserID string `bson:"relatedUserID" json:"relatedUserID"`
}