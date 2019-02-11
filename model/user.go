package model

import "github.com/go-bongo/bongo"

// User is the model for attendees and DJs
type User struct {
	bongo.DocumentBase `bson:",inline"`
	FirstName          string   `json: "firstname"` // First name of user
	LastName           string   `json: "lastname"`  // Last name of user
	Roles              []string `json: "roles"`     // Roles the user has
	PhoneNum           string   `json: "phonenum"`  // Phone number of user
	Code               string   `json: "code"`      // Code for phone number validation
}
