package main

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

var config = Config{
	Database{
		Username: "audiocache",
		Password: "audiocache",
		Hostname: "localhost",
		Database: "audiocache",
		Adapter:  "postgres",
		SSLMode:  "disable",
	},
	Storage{
		Location: "/tmp/audiocache/",
	},
	Server{
		Listen: "127.0.0.1",
		Port:   "8080",
	},
	API{
		Location: "http://localhost:8080/",
		Files:    "files/",
	},
}
