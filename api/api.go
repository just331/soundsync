package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joshuaj1397/soundsync/model"
)

func CreateParty(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	phoneNum := (params["phoneNum"])
	name := (params["name"])
	party, err := model.CreateParty(phoneNum, name)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(party)
}
