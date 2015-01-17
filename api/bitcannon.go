package main

import (
	"fmt"
	// "log"
	"github.com/antonholmquist/jason"
	"io/ioutil"
	"os"
)

var torrentDB *TorrentDB
var err error

const resultLimit int = 100

func main() {
	// Get mongo url from config.json, otherwise default to 127.0.0.1
	mongo := "127.0.0.1"
	bitcannonPort := "1337"
	f, err := ioutil.ReadFile("config.json")
	if err == nil {
		json, err := jason.NewObjectFromBytes(f)
		if err == nil {
			val, err := json.GetString("mongo")
			if err == nil {
				mongo = val
			}
			val, err = json.GetString("bitcannonPort")
			if err == nil {
				bitcannonPort = val
			}
		}
	}
	// Try to connect to the database
	fmt.Println("Connecting to Mongo at " + mongo)
	torrentDB, err = NewTorrentDB(mongo)
	if err != nil {
		fmt.Println("Couldn't connect to Mongo. Please make sure it is installed and running.")
		return
	}
	defer torrentDB.Close()

	if len(os.Args) > 1 {
		importFile(os.Args[1])
	} else {
		runServer(bitcannonPort)
	}
}

func runServer(bitcannonPort string) {
	fmt.Println("Starting the API server.")
	fmt.Println("BitCannon now running on http://127.0.0.1:" + bitcannonPort + "/")
	api := NewAPI()
	api.AddRoutes()
	api.Run(":" + bitcannonPort)
}
