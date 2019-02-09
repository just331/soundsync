package model

import "github.com/go-bongo/bongo"

type User struct {
	bongo.DocumentBase `bson:",inline"`
	FirstName          string
	LastName           string
	Roles              []string
	PhoneNum           string
}
