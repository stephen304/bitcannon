package main

import (
	"bufio"
	"errors"
	"fmt"
	// "log"
	"github.com/antonholmquist/jason"
	"io/ioutil"
	"os"
	"strings"
)

var torrentDB *TorrentDB
var err error

const resultLimit int = 100

func main() {
	// Get mongo url from config.json, otherwise default to 127.0.0.1
	mongo := "127.0.0.1"
	f, err := ioutil.ReadFile("config.json")
	if err == nil {
		json, err := jason.NewObjectFromBytes(f)
		if err == nil {
			val, err := json.GetString("mongo")
			if err == nil {
				mongo = val
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
		runServer()
	}
}

func runServer() {
	fmt.Println("Starting the API server.")
	fmt.Println("BitCannon now running on http://127.0.0.1:" + "1337" + "/")
	api := NewAPI()
	api.AddRoutes()
	api.Run(":1337")
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
	imported := 0
	skipped := 0
	for scanner.Scan() {
		status, _ := importLine(scanner.Text())
		if status {
			imported++
		} else {
			skipped++
		}
	}

	fmt.Println("File parsing ended.")
	fmt.Print("New: ")
	fmt.Println(imported)
	fmt.Print("Dup: ")
	fmt.Println(skipped)

	if err := scanner.Err(); err != nil {
		fmt.Println("The program encountered an error while reading the file. Ensure that the file isn't corrupted.")
		return
	}
}

func importLine(line string) (bool, error) {
	if strings.Count(line, "|") != 4 {
		return false, errors.New("Something's up with this torrent. Expected 5 values separated by |.")
	}
	data := strings.Split(line, "|")
	return torrentDB.Insert(data[0], data[1], data[2], data[3], data[4])
}
