package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"CacheIndex",
		"GET",
		"/caches",
		CacheIndex,
	},
	Route{
		"CacheShow",
		"GET",
		"/caches/{cacheId}",
		CacheShow,
	},
	Route{
		"CacheCreate",
		"POST",
		"/caches",
		CacheCreate,
	},
}
