package model

// User is the model for attendees and DJs
type User struct {
	NickName string   `json: "nickName" bson: "nickName"` // Used to uniquely identify someone in a party instance
	Roles    []string `json: "roles" bson: "roles"`       // Roles the user has
	PhoneNum string   `json: "phoneNum" bson: "phoneNum"` // Phone number of user
	Code     string   `json: "code" bson: "code"`         // Code for phone number validation and logging in
}
