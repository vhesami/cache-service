package main

import "net/http"

type Route struct {
	Name            string
	Path            string
	Methods         string
	HandlerFunction http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		Name:            "Index",
		Path:            "/",
		Methods:         "GET",
		HandlerFunction: Index,
	},
	Route{
		Name:            "ProcessQuery",
		Path:            "/query",
		Methods:         "POST",
		HandlerFunction: StoreQuery,
	},
}
