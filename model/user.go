package model

import "github.com/go-bongo/bongo"

// User is the model for attendees and DJs
type User struct {
	bongo.DocumentBase `bson:",inline"`
	FirstName          string   // First name of user
	LastName           string   // Last name of user
	Roles              []string // Roles the user has
	PhoneNum           string   // Phone number of user
	Code               string   // Code for phone number validation
}
