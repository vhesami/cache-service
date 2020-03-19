package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Listening to :8080")
	router := InitializeRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
