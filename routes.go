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
		Path:            "/store",
		Method:          "POST",
		HandlerFunction: StoreHandler,
	},
	Route{
		Name:            "RetrieveCache",
		Path:            "/retrieve",
		Method:          "POST",
		HandlerFunction: RetrieveHandler,
	},
}
