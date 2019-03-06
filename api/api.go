package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joshuaj1397/soundsync/model"
)

var codeLength = 6

func CreateParty(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	phoneNum := params["phoneNum"]
	name := params["name"]

  // Generate random code
	rand.Seed(time.Now().UnixNano())
	const letterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	partyCode := make([]byte, codeLength)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	party, err := model.CreateParty(phoneNum, name, partyCode)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(party)
}

func JoinParty(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	partyCode := (arams["partyCode"]
	name := params["name"]
	party, err := model.JoinParty(partyCode, name)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(party)
}

func Verify(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	phoneNum := params["phoneNum"]
	name := params["name"]
  authCode := params["authCode"]
	status, err := model.Verify(phoneNum, name, authCode)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(status)
}

func LinkSpotify(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
  //TODO: Make call to spotify API
	json.NewEncoder(w).Encode(spotifyUserId)
}

func SearchSpotify(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query := params["query"]
  //TODO: Make call to spotify API
	json.NewEncoder(w).Encode(songs)
}

func AddSong(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	songId := params["songId"]
  partyId := params["partyId"]
  status, err := model.AddSong(songId, partyId)
  if err != nil {
    log.Fatal(err)
  }
	json.NewEncoder(w).Encode(status)
}

func SongQueue(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
  partyId := params["partyId"]
  songQueue, err := model.SongQueue(partyId)
  if err != nil {
    log.Fatal(err)
  }
	json.NewEncoder(w).Encode(songQueue)
}

func SkipSong(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	songId := params["songId"]
  partyId := params["partyId"]
  //TODO: Make call to spotify API
  status, err := model.SkipSong(songId, partyId)
  if err != nil {
    log.Fatal(err)
  }
	json.NewEncoder(w).Encode(status)
}

func RemoveSong(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	songId := params["songId"]
  partyId := params["partyId"]
  status, err := model.RemoveSong(songId, partyId)
  if err != nil {
    log.Fatal(err)
  }
	json.NewEncoder(w).Encode(status)
}
