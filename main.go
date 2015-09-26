package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var config Config

func main() {
	var err error
	config, err = LoadFile("config.json")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error loading config")
		os.Exit(1)
	}
	router := NewRouter()
	log.Fatal(http.ListenAndServe(config.Server.Listen+":"+config.Server.Port, router))
}
