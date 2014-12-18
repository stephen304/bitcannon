package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"os"
	"bufio"
)

var session *mgo.Session
var err error

func main() {
	// Try to connect to the database
	session, err = mgo.Dial("127.0.0.1")
	if err != nil {
		fmt.Println("Couldn't connect to Mongo. Please make sure it is installed and running.")
		return
	}

	if len(os.Args) > 1 {
		importFile(os.Args[1])
	} else {
		//Run server?
	}
}

func importFile(filename string) {
	fmt.Print("Attempting to parse ")
	fmt.Println(filename)

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening the file. Make sure it exists and is readable.")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("The program encountered an error while reading the file. Ensure that the file isn't corrupted")
		return
	}
}

// hash, title, category, details, download /*hash string, title string, category string, details string*/
func importHash() {
	_ = session.DB("bitcannon").C("torrents")
}
