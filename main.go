package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	port := os.Getenv("PORT")
	esUrl := os.Getenv("ES_URL")

	fmt.Printf("ENV : {\r\n\tPORT:%s,\r\n\tES_URL:%s\r\n}\r\n", port, esUrl)

	router := NewMuxRouter()
	InitializeElasticSearchEnvironment(esUrl)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}
