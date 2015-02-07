package main

import ()

func runScheduler() {
	go runAutoUpdate()
	go scrapeWorker()
}
