package main

import ()

func runScheduler() {
	go importScheduler()
	if config.ScrapeEnabled {
		go scrapeWorker()
	}
}
