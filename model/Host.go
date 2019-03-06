package model

// Host is the instance of a party and account for host of the party
type Host struct {
	SpotifyAuth string   `json: "spotifyAuth" bson: "spotifyAuth"` // Auth token for spotify
	PhoneNum    string   `json: "phoneNum" bson: "phonenum"`       // Phone number of host
	Users       []string `json: "users" bson: "users"`             // Other users in this instance
}
