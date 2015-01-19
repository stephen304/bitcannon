package main

import (
	"github.com/Stephen304/goscrape"
	"github.com/oleiade/lane"
)

var scrapeQueue *lane.PQueue = lane.NewPQueue(lane.MINPQ)

func queueBtih(btih string, priority int) {
	scrapeQueue.Push(btih, priority)
}

func scrapeWork() {
	for {
		btihInterface, _ := scrapeQueue.Pop()
		if btihInterface != nil {
			if btih, ok := btihInterface.(string); ok {
				seeders, leechers := multiScrape(btih, trackers)
				torrentDB.Update(btih, seeders, leechers)
			}
		}
	}
}

func multiScrape(btih string, urls []string) (int, int) {
	seed := 0
	leech := 0
	for _, url := range urls {
		newSeed, newLeech, _, err := goscrape.Udp(btih, url)
		if err != nil && newSeed > seed && newLeech > leech {
			seed = newSeed
			leech = newLeech
		}
	}
	return seed, leech
}
