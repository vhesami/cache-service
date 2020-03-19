package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func InitializeRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {

		var httpHandler http.Handler
		httpHandler = route.HandlerFunction
		httpHandler = Logger(httpHandler, route.Name)

		router.
			Methods(route.Methods).
			Path(route.Path).
			Name(route.Name).
			Handler(httpHandler)
	}
	return router
}
