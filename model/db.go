package model

import (
	"log"
	"os"

	"github.com/globalsign/mgo/bson"
	"github.com/go-bongo/bongo"
)

var (
	config = &bongo.Config{
		ConnectionString: "",
		Database:         "soundsync",
	}

	connection, connErr = bongo.Connect(config)
)

func init() {
	config.ConnectionString = os.Getenv("DefaultConnection")

	connection, connErr = bongo.Connect(config)
	if connErr != nil {
		log.Fatal(connErr)
	}
}

// SaveHost handles creating and updating a host
func SaveHost(myHost *Host) error {
	err := connection.Collection("host").Save(myHost)
	if vErr, ok := err.(*bongo.ValidationError); ok {
		log.Fatal(vErr.Errors)
	} else {
		log.Fatal(err.Error())
	}
	return err
}

// FindHostByID returns a Host given the Id
func FindHostByID(id string) (*Host, error) {
	host := &Host{}
	err := connection.Collection("host").FindById(bson.ObjectIdHex(id), host)
	return host, err
}

// FindHostByPhoneNum returns a Host given the phone number
func FindHostByPhoneNum(phoneNum string) (*Host, error) {
	host := &Host{}
	err := connection.Collection("host").FindOne(bson.M{"PhoneNum:": phoneNum}, host)
	return host, err
}

// DeleteHost deletes a Host given host instance
func DeleteHost(myHost *Host) error {
	err := connection.Collection("host").DeleteDocument(myHost)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

// DeleteHostByID deletes a Host given a Host id
func DeleteHostByID(id string) error {
	host, err := FindHostByID(id)
	if err != nil {
		return err
	}

	err = DeleteHost(host)
	return err
}

// DeleteHostByPhoneNum deletes a Host given a Host phone number
func DeleteHostByPhoneNum(phoneNum string) error {
	host, err := FindHostByPhoneNum(phoneNum)
	if err != nil {
		return err
	}

	err = DeleteHost(host)
	return err
}
