package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// Health check api
// Returns response when app is running
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "IP: %s and Time: %s\r\n", r.RemoteAddr, GetCurrentLocalTime().Format("2006-01-02 15:04:05.000"))
}

// This rest api provides storing user queries in cache
// Input: user_id and text
// Output: status and number of stored tokens
func StoreHandler(w http.ResponseWriter, r *http.Request) {
	var request StoreRequest
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Panicf("StoreHandler() ERROR: %v\r\n", err)
	}
	if err := r.Body.Close(); err != nil {
		log.Panicf("StoreHandler() ERROR: %v\r\n", err)
	}
	if err := json.Unmarshal(body, &request); err != nil {
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Printf("StoreHandler() ERROR: %v\r\n", err)
		}
	}

	client := GetElasticClient()
	tokens := StoreCache(client, request)
	response := StoreResponse{Success: true, StoredTokenCount: len(tokens)}
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("StoreHandler() ERROR: %v\r\n", err)
	}
}

// This rest api provides frequent tokens of specific user in last defined period in csv format
// Input: user_id,recency hours, result size, tokens type and delimiter
// Output: delimiter separated values by recency and frequency
func RetrieveHandler(w http.ResponseWriter, r *http.Request) {
	var request RetrieveRequest
	w.Header().Set("Content-Type", "plain/text; charset=UTF-8")

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Panicf("RetrieveHandler() ERROR: %v\r\n", err)
	}
	if err := r.Body.Close(); err != nil {
		log.Panicf("RetrieveHandler() ERROR: %v\r\n", err)
	}
	if err := json.Unmarshal(body, &request); err != nil {
		w.WriteHeader(422) // unprocessable entity
	}

	client := GetElasticClient()
	csv := RetrieveCache(client, request)

	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintln(w, csv)
}
