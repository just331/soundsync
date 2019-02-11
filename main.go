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
	// Routes
	http.HandleFunc("/hostparty", HostParty)

	// I'll be right by your side till 3005
	log.Fatal(http.ListenAndServe(":"+port, nil))
	fmt.Println("Listening on port: " + port)
}

func HostParty(w http.ResponseWriter, r *http.Request) {
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
