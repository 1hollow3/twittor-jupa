package db

import (
	"context"
	"github.com/1hollow3/twittor-jupa/models"
	"go.mongodb.org/mongo-driver/bson"
	options2 "go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func ReadAllUsers(ID string, page int64, search string, typeOf string) ([]*models.User, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("usernames")

	var result []*models.User

	options := options2.Find()
	options.SetSkip((page - 1) * 20)
	options.SetLimit(20)

	query := bson.M{
		"name": bson.M{"$regex": `(?i)` + search},
	}

	pointer, err := col.Find(ctx, query, options)
	if err != nil {
		log.Fatal(err.Error())
		return result, false
	}

	var found, include bool

	for pointer.Next(ctx) {
		var user models.User
		err := pointer.Decode(&user)
		if err != nil {
			return result, false
		}

		var r models.Relation
		r.UserID = ID
		r.RelatedUserID = user.ID.Hex()

		include = false

		found, err = ConsultRelation(r)
		if typeOf == "new" || (typeOf == "follow" && found == true) {
			include = true
		}
		if r.UserID == r.RelatedUserID {
			include = false
		}

		if include {
			user.Password = ""
			user.Bio = ""
			user.WebSite = ""
			user.Location = ""
			user.Banner = ""
			user.Email = ""

			result = append(result, &user)
		}
	}

	err = pointer.Err()
	if err != nil {
		return result, false
	}

	pointer.Close(ctx)
	return result, true
}
