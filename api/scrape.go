package main

import (
	"fmt"
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
			fmt.Println("no work pause for 1 sec")
			time.Sleep(5 * time.Second)
		}
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
