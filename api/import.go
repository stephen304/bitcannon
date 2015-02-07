package main

import (
	"bufio"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
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
	fmt.Print("[OK!] Finished auto importing.")
}

func importWorker(url string, freq int) {
	for _ = range time.Tick(time.Duration(freq) * time.Second) {
		importURL(url)
	}
}

func importFile(filename string) {
	// Print out status
	fmt.Print("[OK!] Attempting to parse ")
	fmt.Println(filename)

	// Try to open the file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("[ERR] Sorry! I Couldn't access the specified file.")
		fmt.Println("      Double check the permissions and file path.")
		return
	}
	defer file.Close()
	fmt.Println("[OK!] File opened")

	// Check file extension
	var gzipped bool = false
	if strings.HasSuffix(filename, ".txt") {
		gzipped = false
	} else if strings.HasSuffix(filename, ".csv") {
		gzipped = false
	} else if strings.HasSuffix(filename, ".txt.gz") {
		gzipped = true
	} else {
		fmt.Println("[ERR] My deepest apologies! The file doesn't meet the requirements.")
		fmt.Println("      BitCannon currently accepts .txt and gzipped .txt files only.")
		return
	}
	fmt.Println("[OK!] Extension is valid")
	importReader(file, gzipped)
}

func importURL(url string) {
	fmt.Println("[OK!] Starting to import from url:")
	fmt.Println("      " + url)
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("[ERR] Oh no! Couldn't request torrent updates.")
		fmt.Println("      Is your internet working? Is BitCannon firewalled?.")
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
		fmt.Println("[!!!] I was given a URL that doesn't end in .txt or .txt.gz.")
		fmt.Println("      I'll assume it's regular text.")
	}
	fmt.Println("[OK!] Compression detection complete")
	importReader(response.Body, gzipped)
}

func importReader(reader io.Reader, gzipped bool) {
	var scanner *bufio.Scanner
	if gzipped {
		gReader, err := gzip.NewReader(reader)
		if err != nil {
			fmt.Println("[ERR] My bad! I tried to start uncompressing your archive but failed.")
			fmt.Println("      Try checking the file, or send me the file so I can check it out.")
			return
		}
		defer gReader.Close()
		fmt.Println("[OK!] GZip detected, unzipping enabled")
		scanner = bufio.NewScanner(gReader)
	} else {
		scanner = bufio.NewScanner(reader)
	}
	fmt.Println("[OK!] Reading initialized")
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
	fmt.Println("[OK!] Reading completed")
	fmt.Println("      " + strconv.Itoa(imported) + " torrents imported")
	fmt.Println("      " + strconv.Itoa(skipped) + " torrents skipped")
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
		return torrentDB.Insert(data[0], data[1], data[2], 0, data[3])
	} else if strings.Count(line, "|") == 6 {
		data := strings.Split(line, "|")
		if len(data[2]) != 40 {
			return false, errors.New("Probably not a torrent archive")
		}
		if data[4] == "" {
			data[4] = "Other"
		}
		size, _ = strconv.Atoi(data[1])
		return torrentDB.Insert(data[2], data[0], data[4], size, "")
	} else {
		return false, errors.New("Something's up with this torrent.")
	}
}
