package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/just331/soundsync/api"
	"github.com/just331/soundsync/model"
)

var (
	port         = "3005"
	mySigningKey = []byte("ASuperSecretSigningKeyCreatedByTheAliensFromArrival")
)

var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})

func main() {
	router := mux.NewRouter()
	model.CreateDatabaseClient()

	// API
	router.Handle("/GetToken", api.GetToken).Methods("GET")
	router.Handle("/CreateParty/{nickname}/{phoneNum}/{partyName}", jwtMiddleware.Handler(api.CreateParty)).Methods("POST")
	router.Handle("/JoinParty/{nickname}/{partyCode}/{phoneNum}", jwtMiddleware.Handler(api.JoinParty)).Methods("POST")

	//TODO: Find out what this endpoint needs and returns
	// router.HandleFunc("/LinkSpotify/").Methods("POST")
	// router.HandleFunc("/SearchSpotify/{query}", api.SearchSpotify).Methods("GET")
	// router.HandleFunc("/AddSong/{songId}/{partyId}", api.AddSong).Methods("POST")
	// router.HandleFunc("/SongQueue/{partyId}", api.SongQueue).Methods("GET")
	// router.HandleFunc("/SkipSong/{songId}/{partyId}", api.SkipSong).Methods("POST")
	// router.HandleFunc("/RemoveSong/{songId}/{partyId}", api.RemoveSong).Methods("POST")

	go func() {
		fmt.Println("Listening on port: " + port)
		if err := http.ListenAndServe(":"+port, router); err != nil {
			log.Fatal(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	log.Println("Shutting down soundsync...")
	os.Exit(0)
}
