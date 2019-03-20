package model

import "github.com/globalsign/mgo/bson"

// User is the model for attendees and DJs
type User struct {
	ID       bson.ObjectId `bson: "_id, omitempty"`            // Id for the object on MongoDB
	NickName string        `json: "nickName" bson: "nickName"` // Used to uniquely identify someone in a party instance
	Roles    []string      `json: "roles" bson: "roles"`       // Roles the user has
	PhoneNum string        `json: "phoneNum" bson: "phoneNum"` // Phone number of user
	Code     string        `json: "code" bson: "code"`         // Code for phone number validation and logging in
	Verified bool          `json: "verified" bson: "verified"` // Used to auth the user
}
