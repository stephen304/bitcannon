package main

import (
	"errors"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"time"
)

type TorrentDB struct {
	session    *mgo.Session
	collection *mgo.Collection
}

type Torrent struct {
	Btih     string `bson:"_id,omitempty"`
	Title    string
	Category string
	Size     int
	Details  []string
	Swarm    Stats
	Lastmod  time.Time
	Imported time.Time
}

type Stats struct {
	Seeders  int
	Leechers int
}

func NewTorrentDB(url string) (*TorrentDB, error) {
	session, err := mgo.Dial(url)
	if err != nil {
		return nil, err
	}
	session.SetMode(mgo.Monotonic, true)
	collection := session.DB("bitcannon").C("torrents")
	collection.EnsureIndex(mgo.Index{Key: []string{"$text:title"}, Name: "title"})
	collection.EnsureIndex(mgo.Index{Key: []string{"category"}, Name: "category"})
	collection.EnsureIndex(mgo.Index{Key: []string{"swarm.seeders"}, Name: "seeders"})
	collection.EnsureIndex(mgo.Index{Key: []string{"lastmod"}, Name: "lastmod"})
	return &TorrentDB{session, collection}, nil
}

func (torrentDB *TorrentDB) Close() {
	torrentDB.session.Close()
}

func (torrentDB *TorrentDB) Stats(r render.Render) {
	count, err := torrentDB.collection.Count()
	if err != nil {
		r.JSON(500, map[string]interface{}{"message": "API Error"})
		return
	}
	r.JSON(200, map[string]interface{}{"Count": count, "Trackers": trackers})
}

func (torrentDB *TorrentDB) Categories(r render.Render) {
	var result []string
	err := torrentDB.collection.Find(nil).Distinct("category", &result)
	if err != nil {
		r.JSON(500, map[string]interface{}{"message": "API Error"})
		return
	}
	var size int
	for size = range result {
	}
	stats := make([]map[string]interface{}, size+1, size+1)
	for i, cat := range result {
		total, err := torrentDB.collection.Find(bson.M{"category": cat}).Count()
		if err != nil {
			stats[i] = map[string]interface{}{cat: 0}
		} else {
			stats[i] = map[string]interface{}{"name": cat, "count": total}
		}
	}
	r.JSON(200, stats)
}

func (torrentDB *TorrentDB) Browse(r render.Render, params martini.Params) {
	result := []Torrent{}
	err = torrentDB.collection.Find(bson.M{"category": params["category"]}).Sort("-swarm.seeders").Limit(resultLimit).All(&result)
	if err != nil {
		r.JSON(404, map[string]interface{}{"message": err.Error()})
		return
	}
	r.JSON(200, result)
}

func (torrentDB *TorrentDB) Search(r render.Render, params martini.Params) {
	result := []Torrent{}
	skip := 0
	if value, ok := params["skip"]; ok {
		skip, err = strconv.Atoi(value)
		if err != nil {
			r.JSON(400, map[string]interface{}{"message": err.Error()})
			return
		}
	}
	var pipe *mgo.Pipe
	if category, ok := params["category"]; ok {
		pipe = torrentDB.collection.Pipe([]bson.M{
			{"$match": bson.M{"$text": bson.M{"$search": params["query"]}}},
			{"$match": bson.M{"category": category}},
			{"$sort": bson.M{"swarm.seeders": -1}},
			{"$skip": skip},
			{"$limit": resultLimit},
		})
	} else {
		pipe = torrentDB.collection.Pipe([]bson.M{
			{"$match": bson.M{"$text": bson.M{"$search": params["query"]}}},
			{"$sort": bson.M{"swarm.seeders": -1}},
			{"$skip": skip},
			{"$limit": resultLimit},
		})
	}
	iter := pipe.Iter()
	err = iter.All(&result)
	if err != nil {
		r.JSON(404, map[string]interface{}{"message": err.Error()})
		return
	}
	r.JSON(200, result)
}

func (torrentDB *TorrentDB) Get(r render.Render, params martini.Params) {
	result := Torrent{}
	err = torrentDB.collection.Find(bson.M{"_id": params["btih"]}).One(&result)
	if err != nil {
		r.JSON(404, map[string]interface{}{"message": "Torrent not found."})
		return
	}
	r.JSON(200, result)
}

func (torrentDB *TorrentDB) Insert(btih string, title string, category string, size int, details string) (bool, error) {
	var detailsArr []string
	if details != "" {
		detailsArr = []string{details}
	}
	err := torrentDB.collection.Insert(
		&Torrent{Btih: btih,
			Title:    title,
			Category: category,
			Size:     size,
			Details:  detailsArr,
			Swarm:    Stats{Seeders: -1, Leechers: -1},
			Lastmod:  time.Now(),
			Imported: time.Now(),
		})
	if err != nil {
		return false, errors.New("Something went wrong when trying to insert.")
	}
	return true, nil
}

func (torrentDB *TorrentDB) Update(btih string, seeders int, leechers int) {
	match := bson.M{"_id": btih}
	update := bson.M{"$set": bson.M{"swarm": &Stats{Seeders: seeders, Leechers: leechers}, "lastmod": time.Now()}}
	torrentDB.collection.Update(match, update)
}

func (torrentDB *TorrentDB) GetStale() []string {
	result := []Torrent{}
	err = torrentDB.collection.Find(bson.M{"swarm.seeders": -1, "swarm.leechers": -1}).Limit(50).All(&result)
	if len(result) == 0 {
		// No unscraped torrents, get stale ones
		torrentDB.collection.Find(bson.M{"lastmod": bson.M{"$lt": time.Now().Add(-24 * time.Hour)}}).Sort("lastmod").Limit(50).All(&result)
	}
	var btih = make([]string, len(result))
	for i := range result {
		btih[i] = result[i].Btih
	}
	return btih
}
