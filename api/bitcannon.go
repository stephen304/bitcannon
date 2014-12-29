package main

import (
	"bufio"
	"errors"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	// "log"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/cors"
	"github.com/martini-contrib/render"
	"os"
	"strings"
)

var session *mgo.Session
var collection *mgo.Collection
var err error

type Torrent struct {
	Btih     string `bson:"_id,omitempty"`
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
		return
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	collection = session.DB("bitcannon").C("torrents")
	collection.EnsureIndex(mgo.Index{Key: []string{"$text:title"}, Name: "title"})

	if len(os.Args) > 1 {
		importFile(os.Args[1])
	} else {
		runServer()
	}
}

func runServer() {
	fmt.Println("Starting the API server.")
	m := martini.Classic()
	m.Use(render.Renderer())
	m.Use(cors.Allow(&cors.Options{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"POST", "GET"},
		ExposeHeaders: []string{"Content-Length"},
	}))
	m.Get("/stats", func(r render.Render) {
		count, err := collection.Count()
		if err != nil {
			r.JSON(500, map[string]interface{}{"count": "API Error"})
			return
		}
		r.JSON(200, map[string]interface{}{"count": count})
	})
	m.Get("/browse", func(r render.Render) {
		var result []string
		err := collection.Find(nil).Distinct("category", &result)
		if err != nil {
			r.JSON(500, map[string]interface{}{"error": "API Error"})
			return
		}
		var size int
		for size = range result {
		}
		stats := make([]map[string]interface{}, size+1, size+1)
		for i, cat := range result {
			total, err := collection.Find(bson.M{"category": cat}).Count()
			if err != nil {
				stats[i] = map[string]interface{}{cat: 0}
			} else {
				stats[i] = map[string]interface{}{"name": cat, "count": total}
			}
		}
		r.JSON(200, stats)
	})
	m.Get("/torrent/:btih", func(r render.Render, params martini.Params) {
		result := Torrent{}
		err = collection.Find(bson.M{"btih": params["btih"]}).One(&result)
		if err != nil {
			r.JSON(404, map[string]interface{}{"message": "Torrent not found."})
			return
		}
		r.JSON(200, result)
	})
	m.Get("/search/:query", func(r render.Render, params martini.Params) {
		result := []Torrent{}
		pipe := collection.Pipe([]bson.M{
			{"$match": bson.M{"$text": bson.M{"$search": params["query"]}}},
			{"$sort": bson.M{"score": bson.M{"$meta": "textScore"}}},
		})
		iter := pipe.Iter()
		err = iter.All(&result)
		if err != nil {
			r.JSON(404, map[string]interface{}{"message": err.Error()})
			return
		}
		r.JSON(200, result)
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
