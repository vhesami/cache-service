package main

import (
	"context"
	"github.com/olivere/elastic"
)

func do() {
	client, err := elastic.NewClient(elastic.SetURL("http://msvc.ir:9200"))
	if err != nil {
		// Handle error
	}
	exists, err := client.IndexExists("twitter").Do(context.Background())
	if err != nil {
		// Handle error
	}
	if !exists {
		// Index does not exist yet.
	}
}
