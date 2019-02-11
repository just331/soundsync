package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/joshuaj1397/soundsync/model"
)

func main() {
	// Routes
	http.HandleFunc("/signin", CreateParty)

	// I'll be right by your side till 3005
	log.Fatal(http.ListenAndServe(":3005", nil))
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
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
}
