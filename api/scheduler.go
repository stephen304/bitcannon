package main

import ()

func runScheduler() {
	go importScheduler()
	go scrapeWorker()
}
