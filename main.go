package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	/*
		port := os.Getenv("PORT")
		esUrl := os.Getenv("ES_URL")

		fmt.Printf("ENV : {\r\n\tPORT:%s,\r\n\tES_URL:%s\r\n}\r\n", port, esUrl)

		router := NewMuxRouter()
		InitializeElasticSearchEnvironment(esUrl)

		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
	*/
	router := NewMuxRouter()
	fmt.Println("Listening to 0.0.0.0:8080 ...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
