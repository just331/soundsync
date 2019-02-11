package model

import "github.com/go-bongo/bongo"

// Host is the instance of a party and account for host of the party
type Host struct {
	bongo.DocumentBase `bson:",inline"`
	SpotifyAuth        string   `json: "spotifyauth"` // Auth token for spotify
	PhoneNum           string   `json: "phonenum"`    // Phone number of host
	Users              []string `json: "users"`       // Other users in this instance
}
