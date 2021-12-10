package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoCN = ConnectDB()
var clientOptions = options.Client().ApplyURI("mongodb+srv://juapau:jupaomero1995@twittor.mcglm.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

func ConnectDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Connection Success")
	return client
}

func CheckConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
