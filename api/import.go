package main

import (
	"bufio"
	"compress/gzip"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
	_ "crypto/sha512"
)

func importScheduler() {
	for _, site := range archives {
		url, err := site.GetString("url")
		if err == nil {
			freq, err := site.GetInt64("frequency")
			if err == nil {
				go importWorker(url, int(freq))
			}
			importURL(url)
		}
	}
	log.Print("[OK!] Finished auto importing.")
}

func importWorker(url string, freq int) {
	for _ = range time.Tick(time.Duration(freq) * time.Second) {
		importURL(url)
	}
}

func importFile(filename string) {
	// Print out status
	log.Print("[OK!] Attempting to parse ")
	log.Println(filename)

	// Try to open the file
	file, err := os.Open(filename)
	if err != nil {
		log.Println("[ERR] Sorry! I Couldn't access the specified file.")
		log.Println("      Double check the permissions and file path.")
		return
	}
	defer file.Close()
	log.Println("[OK!] File opened")

	// Check file extension
	var gzipped bool = false
	if strings.HasSuffix(filename, ".txt") {
		gzipped = false
	} else if strings.HasSuffix(filename, ".csv") {
		gzipped = false
	} else if strings.HasSuffix(filename, ".txt.gz") {
		gzipped = true
	} else {
		log.Println("[ERR] My deepest apologies! The file doesn't meet the requirements.")
		log.Println("      BitCannon currently accepts .txt and gzipped .txt files only.")
		return
	}
	log.Println("[OK!] Extension is valid")
	importReader(file, gzipped)
}

func importURL(url string) {
	log.Println("[OK!] Starting to import from url:")
	log.Println("      " + url)
	response, err := http.Get(url)
	if err != nil {
		log.Println("[ERR] Oh no! Couldn't request torrent updates.")
		log.Println("      Is your internet working? Is BitCannon firewalled?.")
		return
	}
	defer response.Body.Close()

	var gzipped bool = false
	if strings.HasSuffix(url, ".txt") {
		gzipped = false
	} else if strings.HasSuffix(url, ".csv") {
		gzipped = false
	} else if strings.HasSuffix(url, ".txt.gz") {
		gzipped = true
	} else {
		log.Println("[!!!] I was given a URL that doesn't end in .txt or .txt.gz.")
		log.Println("      I'll assume it's regular text.")
	}
	log.Println("[OK!] Compression detection complete")
	importReader(response.Body, gzipped)
}

func importReader(reader io.Reader, gzipped bool) {
	var scanner *bufio.Scanner
	if gzipped {
		gReader, err := gzip.NewReader(reader)
		if err != nil {
			log.Println("[ERR] My bad! I tried to start uncompressing your archive but failed.")
			log.Println("      Try checking the file, or send me the file so I can check it out.")
			return
		}
		defer gReader.Close()
		log.Println("[OK!] GZip detected, unzipping enabled")
		scanner = bufio.NewScanner(gReader)
	} else {
		scanner = bufio.NewScanner(reader)
	}
	log.Println("[OK!] Reading initialized")
	imported := 0
	skipped := 0
	// Now we scan ୧༼ಠ益ಠ༽୨
	for scanner.Scan() {
		status, _ := importLine(scanner.Text())
		if status {
			imported++
		} else {
			skipped++
		}
	}
	log.Println("[OK!] Reading completed")
	log.Println("      " + strconv.Itoa(imported) + " torrents imported")
	log.Println("      " + strconv.Itoa(skipped) + " torrents skipped")
}

func importLine(line string) (bool, error) {
	var size int
	if strings.Count(line, "|") == 4 {
		data := strings.Split(line, "|")
		if len(data[0]) != 40 {
			return false, errors.New("Probably not a torrent archive")
		}
		if data[2] == "" {
			data[2] = "Other"
		}
		if (!categoryInBlacklisted(data[2],blacklistedCategories)) {
			return torrentDB.Insert(data[0], data[1], data[2], 0, data[3])
		} else {
			return false, errors.New("Skipping torrent due to category "+data[2]+" in blacklist")
		}
	} else if strings.Count(line, "|") == 6 {
		data := strings.Split(line, "|")
		if len(data[2]) != 40 {
			return false, errors.New("Probably not a torrent archive")
		}
		if data[4] == "" {
			data[4] = "Other"
		}
		size, _ = strconv.Atoi(data[1])
		if (!categoryInBlacklisted(data[4],blacklistedCategories)) {
			return torrentDB.Insert(data[2], data[0], data[4], size, "")
		} else {
			return false, errors.New("Skipping torrent due to category "+data[4]+" in blacklist")
		}
	} else {
		return false, errors.New("Something's up with this torrent.")
	}
}

//http://stackoverflow.com/questions/15323767/how-to-if-x-in-array-in-golang
func categoryInBlacklisted(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}
