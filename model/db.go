package model

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/joshuaj1397/gotwilio"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	connectionStr = os.Getenv("connectionStr")
	dbName        = "soundsync"
	ctx, _        = context.WithTimeout(context.Background(), 10*time.Second)
	client        *mongo.Client
	db            *mongo.Database
	codeLength    = 6
)

func init() {
	client, _ = mongo.NewClient(options.Client().ApplyURI(connectionStr))
	connErr := client.Connect(ctx)
	db = client.Database(dbName)
	fmt.Println("Connected to MongoDB")

	if connErr != nil {
		panic(connErr)
	}

	err := client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateUser(phoneNum, nickname, role string) (interface{}, error) {
	var roles []string
	roles = append(roles, role)
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)

	rand.Seed(time.Now().UnixNano())
	const numberBytes = "0123456789"
	userCode := make([]byte, codeLength)
	for i := range userCode {
		userCode[i] = numberBytes[rand.Intn(len(numberBytes))]
	}

	userBson := bson.M{
		"nickName": nickname,
		"phoneNum": phoneNum,
		"role":     roles,
		"code":     string(userCode),
		"verified": false,
	}

	res, err := db.Collection("User").InsertOne(ctx, userBson)
	if err != nil {
		log.Fatal(err)
	}

	gotwilio.SendMsg(phoneNum, "Your user verification code is: "+string(userCode))
	return res.InsertedID, nil
}

func VerifyUser(phoneNum, userCode string) error {
	user := &User{}
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)

	if user.Code == userCode {
		_, err := db.Collection("User").UpdateOne(ctx, bson.M{"phoneNum": phoneNum}, bson.M{"verified": true})
		if err != nil {
			return errors.New("User not found with that phone number")
		}
	}
	return nil
}

func CreateParty(partyName, phoneNum, nickname string) (string, error) {
	var users []primitive.ObjectID
	user := &User{}
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)

	// Generate random code
	// TODO: Verify uniqueness
	rand.Seed(time.Now().UnixNano())
	const letterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	partyCode := make([]byte, codeLength)
	for i := range partyCode {
		partyCode[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	err := db.Collection("User").FindOne(ctx, bson.M{"phoneNum": phoneNum}).Decode(user)
	// We didn't find a user
	if err != nil {
		Id, err := CreateUser(phoneNum, nickname, "host")
		if err != nil {
			log.Fatal(err)
		}

		// Put the Id of the user in the users slice for the party
		users = append(users, Id.(primitive.ObjectID))
	} else {
		users = append(users, user.ID)
	}

	partyBson := bson.M{
		"partyName":   partyName,
		"spotifyAuth": "", // User will add spotify later
		"partyCode":   string(partyCode),
		"users":       users,
	}

	db.Collection("Party").InsertOne(ctx, partyBson)
	gotwilio.SendMsg(phoneNum, "Your party code is: "+string(partyCode))
	return string(partyCode), nil
}
