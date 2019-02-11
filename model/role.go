package model

import "github.com/go-bongo/bongo"

// Role is the privileges users are given.
type Roles struct {
	bongo.DocumentBase `bson:",inline"`
	Spotify
	Party
}

// Spotify privileges for users.
type Spotify struct {
	skip    bool // skip a song
	pause   bool // pause a song
	volume  bool // change the volume w/in spotify
	shuffle bool // shuffle the playlist
	repeat  bool // repeat a song
	loopQ   bool // loop the entire playlist
	queue   bool //	queue a song
}

// Party privileges for users.
type Party struct {
	kill   bool // end the instance
	kick   bool // kick a user (10 mins, can join back)
	invite bool // invite a user
	ban    bool // ban a user from instance
}
