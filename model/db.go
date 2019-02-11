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

// SaveUser handles creating and updating a user
func SaveUser(myUser *User) error {
	err := connection.Collection("user").Save(myUser)
	if vErr, ok := err.(*bongo.ValidationError); ok {
		log.Fatal(vErr.Errors)
	} else {
		log.Fatal(err.Error())
	}
	return err
}

// FindHostByID returns a Host given the Id
func FindUserByID(id string) (*User, error) {
	user := &User{}
	err := connection.Collection("user").FindById(bson.ObjectIdHex(id), user)
	return user, err
}

// FindUserByPhoneNum returns a User given the phone number
func FindUserByPhoneNum(phoneNum string) (*User, error) {
	user := &User{}
	err := connection.Collection("user").FindOne(bson.M{"PhoneNum:": phoneNum}, user)
	return user, err
}

// DeleteUser deletes a User given user instance
func DeleteUser(myUser *User) error {
	err := connection.Collection("user").DeleteDocument(myUser)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

// DeleteUserByID deletes a User given a User id
func DeleteUserByID(id string) error {
	user, err := FindUserByID(id)
	if err != nil {
		return err
	}

	err = DeleteUser(user)
	return err
}

// DeleteUserByPhoneNum deletes a User given a User phone number
func DeleteUserByPhoneNum(phoneNum string) error {
	user, err := FindUserByPhoneNum(phoneNum)
	if err != nil {
		return err
	}

	err = DeleteUser(user)
	return err
}

// SaveRoles handles creating and updating user roles
func SaveRoles(myRoles *Roles) error {
	err := connection.Collection("roles").Save(myRoles)
	if vErr, ok := err.(*bongo.ValidationError); ok {
		log.Fatal(vErr.Errors)
	} else {
		log.Fatal(err.Error())
	}
	return err
}

// FindRoleByID handles IDing roles of user by their ID
func FindRoleByID(id string) (*Roles, error) {
	role := &Roles{}
	err := connection.Collection("role").FindById(bson.ObjectIdHex(id), role)
	return role, err
}

// FindRolesByPhoneNum returns a roles given the phone number of user
func FindRolesByPhoneNum(phoneNum string) (*Roles, error) {
	roles := &Roles{}
	err := connection.Collection("roles").FindOne(bson.M{"PhoneNum:": phoneNum}, roles)
	return roles, err
}

// DeleteRoles deletes a users roles
func DeleteRoles(myRoles *Roles) error {
	err := connection.Collection("role").DeleteDocument(myRoles)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

// DeleteRolesByID deletes roles given a User id
func DeleteRolesByID(id string) error {
	roles, err := FindUserByID(id)
	if err != nil {
		return err
	}

	err = DeleteUser(roles)
	return err
}

// DeleteUserByPhoneNum deletes roles given a user's phone number
func DeleteRolesByPhoneNum(phoneNum string) error {
	roles, err := FindUserByPhoneNum(phoneNum)
	if err != nil {
		return err
	}

	err = DeleteUser(roles)
	return err
}
