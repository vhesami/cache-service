package main

import "net/http"

type Route struct {
	Name            string
	Path            string
	Method          string
	HandlerFunction http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		Name:            "Index",
		Path:            "/",
		Method:          "GET",
		HandlerFunction: IndexHandler,
	},
	Route{
		Name:            "StoreCache",
		Path:            "/cache/api/v1/store",
		Method:          "POST",
		HandlerFunction: StoreHandler,
	},
	Route{
		Name:            "RetrieveCache",
		Path:            "/cache/api/v1/retrieve",
		Method:          "POST",
		HandlerFunction: RetrieveHandler,
	},
}
