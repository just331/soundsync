package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// User is the model for attendees and DJs
type User struct {
	ID       primitive.ObjectID              `json:"id" bson:"_id, omitempty"` // Id for the object on MongoDB
	NickName string                          `json:"nickName" bson:"nickName"` // Used to uniquely identify someone in a party instance
	Roles    map[primitive.ObjectID][]string `json:"roles" bson:"roles"`       // Roles the user has
	PhoneNum string                          `json:"phoneNum" bson:"phoneNum"` // Phone number of user
	Code     string                          `json:"code" bson:"code"`         // Code for phone number validation and logging in
	Verified bool                            `json:"verified" bson:"verified"` // Used to auth the user
}
