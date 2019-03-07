package model

// Party is the instance of a party and account for Party of the party
type Party struct {
	PartyName   string   `json: "partyName" bson: "partyName"`     // Used to identify party for attendees and hosts
	SpotifyAuth string   `json: "spotifyAuth" bson: "spotifyAuth"` // Auth token for spotify
	PartyCode   string   `json: "partyCode" bson: "partyCode"`     // Party code for attendees to join the party
	Users       []string `json: "users" bson: "users"`             // Other users in this instance
}
