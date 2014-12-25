package main

import (
	"bufio"
	"errors"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	// "log"
	"github.com/go-martini/martini"
	"os"
	"strings"
)

var session *mgo.Session
var collection *mgo.Collection
var err error

type Torrent struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	Btih     string
	Title    []string
	Category []string
	Details  []string
	Download []string
}

func main() {
	// Try to connect to the database
	session, err = mgo.Dial("10.0.1.12")
	if err != nil {
		fmt.Println("Couldn't connect to Mongo. Please make sure it is installed and running.")
		//return
	} else {
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		collection = session.DB("bitcannon").C("torrents")
		collection.EnsureIndex(mgo.Index{Key: []string{"$text:title"}, Name: "title"})
	}

	if len(os.Args) > 1 {
		importFile(os.Args[1])
	} else {
		runServer()
	}
}

func runServer() {
	fmt.Println("Starting the API server.")
	m := martini.Classic()
	m.Get("/stats", func() string {
		return `{
        \"ns\" : \"app.users\",             // namespace
        \"count\" : 9,                    // number of documents
        \"size\" : 432,                   // collection size in bytes
        \"avgObjSize\" : 48,              // average object size in bytes
        \"storageSize\" : 3840,           // (pre)allocated space for the collection in bytes
        \"numExtents\" : 1,               // number of extents (contiguously allocated chunks of datafile space)
        \"nindexes\" : 2,                 // number of indexes
        \"lastExtentSize\" : 3840,        // size of the most recently created extent in bytes
        \"paddingFactor\" : 1,            // padding can speed up updates if documents grow
        \"flags\" : 1,
        \"totalIndexSize\" : 16384,       // total index size in bytes
        \"indexSizes\" : {                // size of specific indexes in bytes
                \"_id_\" : 8192,
                \"username\" : 8192
        },
        \"ok\" : 1
}`
	})
	m.Get("/torrent/:btih", func(params martini.Params) string {
		return `{
			"btih": "` + params["btih"] + `",
			"title": "This is a title"
}`
	})
	m.RunOnAddr(":1337")
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
	return importHash(data[0], data[1], data[2], data[3], data[4])
}

// hash, title, category, details, download /*hash string, title string, category string, details string*/
func importHash(btih string, title string, category string, details string, download string) (bool, error) {
	err := collection.Insert(&Torrent{Btih: btih, Title: []string{title}, Category: []string{category}, Details: []string{details}, Download: []string{download}})
	if err != nil {
		return false, errors.New("Something went wrong when trying to insert.")
	}
	return true, nil
}
