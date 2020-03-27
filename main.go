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

	fmt.Printf("ENV : {\n\tPORT:%s,\n\tES_URL:%s\n}\n", port, esUrl)

	router := NewMuxRouter()
	InitializeElasticSearchEnvironment(esUrl)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}
