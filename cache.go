package main

import (
	"time"
)

var FileRoot string = "http://localhost:8080/files/"

type APICache struct {
	Id        uint64    `json:"id"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Created   time.Time `json:"created"`
	URI       string    `json:"uri"`
}

type APICaches []APICache

type DBCache struct {
	Id        uint64
	Latitude  float64
	Longitude float64
	Created   time.Time
	Path      string
}

type DBCaches []DBCache

type PostCache struct {
	Latitude  float64
	Longitude float64
	Data      string
}

func PathToURI(path string) string {
	return FileRoot + path
}

func DBToAPI(db DBCache) APICache {
	var api APICache

	api.Id = db.Id
	api.Latitude = db.Latitude
	api.Longitude = db.Longitude
	api.Created = db.Created
	api.URI = PathToURI(db.Path)

	return api
}

func PostToDB(post PostCache, filename string) DBCache {
	var db DBCache

	db.Latitude = post.Latitude
	db.Longitude = post.Longitude
	db.Created = time.Now()
	db.Path = filename

	return db
}
