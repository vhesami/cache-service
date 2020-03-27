package main

import (
	"fmt"
	"testing"
)

func TestStoreCache(t *testing.T) {
	esUrl := "http://msvc.ir:9200"
	request := StoreRequest{UserId: "1234", Text: "Go how to program?"}

	InitializeElasticSearchEnvironment(esUrl)
	client := GetElasticClient()
	response := StoreCache(client, request)
	if len(response) != 5 {
		t.Errorf("Invalid tokens count: %d", len(response))
	}
	for id, token := range response {
		fmt.Printf("%s --> %s\r\n", id, token.Token)
	}
}
func TestRetrieveCache(t *testing.T) {
	esUrl := "http://msvc.ir:9200"
	request := RetrieveRequest{UserId: "1234"}

	InitializeElasticSearchEnvironment(esUrl)
	client := GetElasticClient()
	response := RetrieveCache(client, request)
	if len(response) == 0 {
		t.Errorf("Invalid result: %s", response)
	}
	fmt.Println(response)
}
