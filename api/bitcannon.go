package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"os"
)

var session *mgo.Session
var err error

func main() {
	// Try to connect to the database
	session, err = mgo.Dial("127.0.0.0")
	if err != nil {
		fmt.Println("Couldn't connect to Mongo. Please make sure it is installed and running.")
		os.Exit(1)
	}

	importHash()
}

func importFile(filename string) {
	//
}

// hash, title, category, details, download /*hash string, title string, category string, details string*/
func importHash() {
	_ = session.DB("bitcannon").C("torrents")
}
