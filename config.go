package main

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Database Database
	Storage  Storage
	Server   Server
	API      API
}

type Database struct {
	Username string
	Password string
	Hostname string
	Database string
	Adapter  string
	SSLMode  string
}

type Storage struct {
	Location string
}

type Server struct {
	Port   string
	Listen string
}

type API struct {
	Location string
	Files    string
}

func LoadFile(filename string) (Config, error) {
	var config Config

	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(file, &config)
	return config, err
}
