package model

import (
	"log"
	"os"

	"github.com/go-bongo/bongo"
)

var config = &bongo.Config{
	ConnectionString: "",
	Database:         "soundsync",
}

func init() {
	config.ConnectionString = os.Getenv("DefaultConnection")

	connection, err := bongo.Connect(config)
	if err != nil {
		log.Fatal(err)
	}

}
