package main

import (
	"github.com/Stephen304/goscrape"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"time"
)

func scrapeWorker() {
	bulk := goscrape.NewBulk(trackers)
	for {
		stale := torrentDB.GetStale()
		if len(stale) > 1 {
			results := bulk.ScrapeBulk(stale)
			multiUpdate(results)
		} else {
			time.Sleep(30 * time.Second)
		}
		time.Sleep(time.Duration(config.ScrapeDelay) * time.Second)
	}
}

func multiUpdate(results []goscrape.Result) {
	for _, result := range results {
		torrentDB.Update(result.Btih, result.Seeders, result.Leechers)
	}
}

func apiScrape(r render.Render, params martini.Params) {
	result := goscrape.Single(trackers, []string{params["btih"]})[0]
	multiUpdate([]goscrape.Result{result})
	r.JSON(200, map[string]interface{}{"Swarm": map[string]interface{}{"Seeders": result.Seeders, "Leechers": result.Leechers}, "Lastmod": time.Now()})
}
