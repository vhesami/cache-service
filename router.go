package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func NewMuxRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {

		var httpHandler http.Handler
		httpHandler = route.HandlerFunction
		httpHandler = logger(httpHandler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Path).
			Name(route.Name).
			Handler(httpHandler)
	}
	return router
}
func logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		inner.ServeHTTP(w, r)
		log.Printf("\t%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start))
	})
}
