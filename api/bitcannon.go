package main

import (
	"bufio"
	"github.com/antonholmquist/jason"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	ScrapeEnabled bool
	ScrapeDelay   int
}

var config = Config{ScrapeEnabled: false, ScrapeDelay: 0}
var trackers []string
var blacklistedCategories []string
var archives []*jason.Object
var torrentDB *TorrentDB
var err error

const resultLimit int = 200

func main() {
	// Get mongo url from config.json, otherwise default to 127.0.0.1
	mongo := "127.0.0.1"
	bitcannonPort := "1337"
	bitcannonBindIp := "0.0.0.0"
	f, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Println("[!!!] Config not loaded")
	} else {
		json, err := jason.NewObjectFromBytes(f)
		if err == nil {
			// Get mongo connection details
			val, err := json.GetString("mongo")
			if err == nil {
				mongo = val
			}
			// Get desired port
			val, err = json.GetString("bitcannonPort")
			if err == nil {
				bitcannonPort = val
			}
			// Get archive sources
			arc, err := json.GetObjectArray("archives")
			if err == nil {
				archives = arc
			}
			// Get trackers
			trac, err := json.GetStringArray("trackers")
			if err == nil {
				trackers = trac
			}
			// Get blacklisted categories
			blackCats, err := json.GetStringArray("blacklisted_categories")
			if err == nil {
				blacklistedCategories = blackCats
			}
			// Get scraping enabled
			scrape, err := json.GetBoolean("scrapeEnabled")
			if err == nil {
				config.ScrapeEnabled = scrape
			}
			// Get scrape delay
			scrapeDelay, err := json.GetInt64("scrapeDelay")
			if err == nil {
				config.ScrapeDelay = int(scrapeDelay)
			}
			// Get desired listening host
			val, err = json.GetString("bitcannonBindIp")
			if err == nil {
				bitcannonBindIp = val
			}
		}
	}
	// Try to connect to the database
	log.Println("[OK!] Connecting to Mongo at " + mongo)
	torrentDB, err = NewTorrentDB(mongo)
	if err != nil {
		log.Println("[ERR] I'm sorry! I Couldn't connect to Mongo.")
		log.Println("      Please make sure it is installed and running.")
		return
	}
	defer torrentDB.Close()

	if len(os.Args) > 1 {
		importFile(os.Args[1])
		enterExit()
	} else {
		runServer(bitcannonPort, bitcannonBindIp)
	}
}

func runServer(bitcannonPort string, bitcannonBindIp string) {
	log.Println("[OK!] BitCannon is live at http://127.0.0.1:" + bitcannonPort + "/")
	api := NewAPI()
	api.AddRoutes()
	runScheduler()
	api.Run(bitcannonBindIp + ":" + bitcannonPort)
}

func enterExit() {
	log.Println("\n\nPress enter to quit...")
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
}
