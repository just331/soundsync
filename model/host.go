package model

import "github.com/go-bongo/bongo"

type Host struct {
	bongo.DocumentBase `bson:",inline"`
	SpotifyAuth        string
	PhoneNum           string
	Code               string
	Users              []string
}
