package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	auth0 "github.com/auth0-community/go-auth0"
	"github.com/gorilla/mux"
	"github.com/joshuaj1397/soundsync/api"
	jose "gopkg.in/square/go-jose.v2"
)

var (
	port          = "3005"
	auth0Domain   = os.Getenv("AUTH0_DOMAIN")
	auth0ClientID = os.Getenv("AUTH0_CLIENT_ID")
	auth0Secret   = os.Getenv("AUTH0_CLIENT_SECRET")
)

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		secretProvider := auth0.NewKeyProvider(auth0Secret)
		audience := []string{"{YOUR-AUTH0-API-AUDIENCE}"}

		configuration := auth0.NewConfiguration(secretProvider, audience, "https://"+auth0Domain+".auth0.com/", jose.HS256)
		validator := auth0.NewValidator(configuration)

		token, err := validator.ValidateRequest(r)

		if err != nil {
			fmt.Println(err)
			fmt.Println("Token is not valid:", token)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func main() {
	router := mux.NewRouter()

	// API
	router.Handle("/GetToken", api.GetToken).Methods("GET")
	router.Handle("/CreateParty/{nickname}/{phoneNum}/{partyName}", authMiddleware(api.CreateParty)).Methods("POST")
	router.Handle("/JoinParty/{nickname}/{partyCode}/{phoneNum}", authMiddleware(api.JoinParty)).Methods("POST")
	// router.HandleFunc("/Verify/{phoneNum}/{name}/{authCode}", api.Verify).Methods("POST")

	//TODO: Find out what this endpoint needs and returns
	// router.HandleFunc("/LinkSpotify/").Methods("POST")
	// router.HandleFunc("/SearchSpotify/{query}", api.SearchSpotify).Methods("GET")
	// router.HandleFunc("/AddSong/{songId}/{partyId}", api.AddSong).Methods("POST")
	// router.HandleFunc("/SongQueue/{partyId}", api.SongQueue).Methods("GET")
	// router.HandleFunc("/SkipSong/{songId}/{partyId}", api.SkipSong).Methods("POST")
	// router.HandleFunc("/RemoveSong/{songId}/{partyId}", api.RemoveSong).Methods("POST")

	log.Fatal(http.ListenAndServe(":"+port, router))
	fmt.Println("Listening on port: " + port)
}
