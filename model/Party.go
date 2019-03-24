package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// Party is the instance of a party and account for Party of the party
type Party struct {
	ID          primitive.ObjectID   `json:"id" bson:"_id,omitempty"`        // Id for the object on MongoDB
	PartyName   string               `json:"partyName" bson:"partyName"`     // Used to identify party for attendees and hosts
	SpotifyAuth string               `json:"spotifyAuth" bson:"spotifyAuth"` // Auth token for spotify
	PartyCode   string               `json:"partyCode" bson:"partyCode"`     // Party code for attendees to join the party
	Users       []primitive.ObjectID `json:"users" bson:"users"`             // Other users in this instance
}
