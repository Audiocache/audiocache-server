package main

import (
	"log"
	"net/http"
)

func main() {
	router := NewRouter()
	log.Fatal(http.ListenAndServe(config.Server.Listen+":"+config.Server.Port, router))
}
