package model

import "github.com/go-bongo/bongo"

// Host is the instance of a party and account for host of the party
type Host struct {
	bongo.DocumentBase `bson:",inline"`
	SpotifyAuth        string   // Auth token for spotify
	PhoneNum           string   // Phone number of host
	Users              []string // Other users in this instance
}
