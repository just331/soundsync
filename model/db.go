package model

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	connectionStr = os.Getenv("connectionStr")
	dbName        = "soundsync"
	ctx, _        = context.WithTimeout(context.Background(), 10*time.Second)
	client        *mongo.Client
	db            *mongo.Database
)

func init() {
	client, _ = mongo.NewClient(options.Client().ApplyURI(connectionStr))
	connErr := client.Connect(ctx)
	db = client.Database(dbName)

	if connErr != nil {
		panic(connErr)
	}

	err := client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateParty(partyName, phoneNum, nickname, partyCode string) (Party, error) {
	var users []string
	user := &User{}
	party := &Party{}
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)

	err := db.Collection("Users").FindOne(ctx, bson.M{"phoneNum": phoneNum}).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}

	users = append(users, user.NickName)

	partyBson := bson.M{
		"partyName":   partyName,
		"spotifyAuth": "", // User will add spotify later
		"partyCode":   partyCode,
		"users":       users,
	}

	db.Collection("Party").InsertOne(ctx, partyBson)
}
