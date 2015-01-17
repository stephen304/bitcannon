package main

import (
	"bufio"
	"compress/gzip"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
	} else if strings.HasSuffix(filename, ".txt.gz") {
		gzipped = true
	} else {
		fmt.Println("[ERR] My deepest apologies! The file doesn't meet the requirements.")
		fmt.Println("      BitCannon currently accepts .txt and gzipped .txt files only.")
		return
	}
	fmt.Println("[OK!] Extension is valid")

	var scanner *bufio.Scanner
	if gzipped {
		reader, err := gzip.NewReader(file)
		if err != nil {
			fmt.Println("[ERR] My bad! I tried to start uncompressing your archive.")
			fmt.Println("      Try checking the file, or send me the file so I can check it out.")
			return
		}
		defer reader.Close()
		scanner = bufio.NewScanner(reader)
		fmt.Println("[OK!] GZip detected, unzipping enabled")
	} else {
		scanner = bufio.NewScanner(file)
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
	if strings.Count(line, "|") != 4 {
		return false, errors.New("Something's up with this torrent. Expected 5 values separated by |.")
	}
	data := strings.Split(line, "|")
	return torrentDB.Insert(data[0], data[1], data[2], data[3], data[4])
}
