package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/joshuaj1397/soundsync/model"
)

var (
	codeLength = 6
	port       = "3005"
)

func main() {
	router := mux.NewRouter()

	// API
	router.HandleFunc("/CreateParty/{phoneNum}/{name}", api.CreateParty).Methods("POST")
	router.HandleFunc("/JoinParty/{name}/{partyCode}", api.JoinParty).Methods("POST")
	router.HandleFunc("/Verify/{phoneNum}/{name}/{authCode}", api.Verify).Methods("POST")

	//TODO: Find out what this endpoint needs and returns
	router.HandleFunc("/LinkSpotify/").Methods("POST")
	router.HandleFunc("/SearchSongs/{songName}", api.SearchSongs).Methods("GET")
	router.HandleFunc("/AddSong/{songId}/{partyId}", api.AddSong).Methods("POST")
	router.HandleFunc("/SongQueue/{partyId}", api.SongQueue).Methods("GET")
	router.HandleFunc("/SkipSong/{songId}/{partyId}", api.SkipSong).Methods("POST")
	router.HandleFunc("/RemoveSong/{songId}/{partyId}", api.RemoveSong).Methods("POST")

	log.Fatal(http.ListenAndServe(":"+port, router))
	fmt.Println("Listening on port: " + port)
}

func CreateParty(w http.ResponseWriter, r *http.Request) {
	host := &model.Host{}
	err := json.NewDecoder(r.Body).Decode(host)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Generate random code
	rand.Seed(time.Now().UnixNano())
	const letterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, codeLength)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

}
