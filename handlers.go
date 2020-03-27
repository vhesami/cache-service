package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "IP: %s and Time: %s", r.RemoteAddr, time.Now().Format("2006-01-02 15:04:05.000"))
}
func StoreHandler(w http.ResponseWriter, r *http.Request) {
	var request StoreRequest
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Panicf("StoreHandler() ERROR: %v", err)
	}
	if err := r.Body.Close(); err != nil {
		log.Panicf("StoreHandler() ERROR: %v", err)
	}
	if err := json.Unmarshal(body, &request); err != nil {
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Printf("StoreHandler() ERROR: %v", err)
		}
	}

	client := GetElasticClient()
	count := StoreCache(client, request)
	response := StoreResponse{Success: true, StoredTokenCount: count}
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("StoreHandler() ERROR: %v", err)
	}
}
func RetrieveHandler(w http.ResponseWriter, r *http.Request) {
	var request RetrieveRequest
	w.Header().Set("Content-Type", "plain/text; charset=UTF-8")

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Panicf("RetrieveHandler() ERROR: %v", err)
	}
	if err := r.Body.Close(); err != nil {
		log.Panicf("RetrieveHandler() ERROR: %v", err)
	}
	if err := json.Unmarshal(body, &request); err != nil {
		w.WriteHeader(422) // unprocessable entity
	}

	client := GetElasticClient()
	csv := RetrieveCache(client, request)

	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintln(w, csv)
}
