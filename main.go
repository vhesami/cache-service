package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	do()
	port := os.Getenv("PORT")
	fmt.Println("Listening to port " + port)

	router := InitializeRouter()
	log.Fatal(http.ListenAndServe(":"+port, router))
}
