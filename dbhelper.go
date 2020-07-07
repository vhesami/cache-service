package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"log"
	"sync"
	"time"
)

var once sync.Once
var instance *elastic.Client

const indexName = "tokens_cache"

func InitializeElasticSearchEnvironment(elasticUrl string) {
	once.Do(func() {
		if instance == nil {
			instance = newElasticClient(elasticUrl)
		}
	})
	createIndex(instance)
}
func GetElasticClient() *elastic.Client {
	return instance
}
func StoreCache(client *elastic.Client, request StoreRequest) map[string]Token {
	tokens := tokenize(client, request)

	var tokensMap map[string]Token
	tokensMap = make(map[string]Token)

	now := GetCurrentLocalTime().Format(time.RFC3339)

	for _, token := range tokens {
		plainDocId := fmt.Sprintf("%s:%s:%t", token.UserId, token.Token, token.IsKeyword)
		documentId := ComputeSHA1(plainDocId)
		token.Occurs = fetchTokenOccurs(client, documentId)
		token.LastUpdate = now
		tokensMap[documentId] = token
	}
	return indexTokens(client, tokensMap)
}
func RetrieveCache(client *elastic.Client, request RetrieveRequest) string {
	request.Compact()
	boolQuery := createFetchQuery(request)
	searchService := client.
		Search().
		Index(indexName).
		Query(boolQuery).
		Size(request.Size).
		Sort("occurs", false)
	searchResult, err := searchService.Do(context.Background())
	if err != nil {
		log.Fatalf("retrieve() ERROR: %v\r\n", err)
	}
	csv := toCsv(searchResult, request.Delimiter)
	return csv
}

//---- InitializeEnvironment subroutines ---------------------------
func newElasticClient(elasticUrl string) *elastic.Client {
	client, err := elastic.NewClient(elastic.SetURL(elasticUrl))
	if err != nil {
		log.Fatalf("NewElasticClient() ERROR: %v\r\n", err)
	}
	return client
}
func createIndex(client *elastic.Client) {
	exists, err := client.IndexExists(indexName).Do(context.Background())
	if err != nil {
		log.Fatalf("createIndex() ERROR: %v\r\n", err)
	}
	if !exists {
		response, err := client.
			CreateIndex(indexName).
			Body(`{"mappings": {
							"properties": {
							  "user_id": {"type": "keyword"},
							  "last_update": {"type": "date"},
							  "token": {"type": "text"},
							  "is_keyword": {"type": "boolean"},
							  "occurs": {"type": "integer"}
							}
						  }}`).
			Do(context.Background())
		if err != nil {
			log.Fatalf("createIndex() ERROR: %v\r\n", err)
		}
		if !response.Acknowledged {
			log.Fatalf("Could not create index '%s'\r\n", indexName)
		}
	}
}

//---- StoreCache subroutines ---------------------------------------
func callTokenizeService(client *elastic.Client, jsonBody string) []string {
	response, err := client.IndexAnalyze().BodyString(jsonBody).Do(context.Background())
	if err != nil {
		log.Panicf("callTokenizeService() ERROR: %v\r\n", err)
	}
	var tokens []string
	tokens = make([]string, len(response.Tokens))
	for i, token := range response.Tokens {
		tokens[i] = token.Token
	}
	return tokens
}
func tokenize(client *elastic.Client, query StoreRequest) []Token {
	bodyTemplate := `{"filter": ["unique"],"tokenizer":"%s","text":"%s"}`

	keywords := callTokenizeService(client, fmt.Sprintf(bodyTemplate, "keyword", query.Text))
	standards := callTokenizeService(client, fmt.Sprintf(bodyTemplate, "standard", query.Text))

	var tokens []Token
	for _, token := range keywords {
		tokens = append(tokens, Token{IsKeyword: true, Token: token, UserId: query.UserId})
	}
	for _, token := range standards {
		tokens = append(tokens, Token{IsKeyword: false, Token: token, UserId: query.UserId})
	}
	return tokens
}
func fetchTokenOccurs(client *elastic.Client, documentId string) *int64 {
	var occurs = new(int64)
	response, err := client.Get().Index(indexName).Id(documentId).FetchSource(false).Do(context.Background())
	if err != nil {
		*occurs = 1
	}
	if response != nil && response.Found {
		occurs = response.Version
		*occurs += 1
	}
	return occurs
}
func indexTokens(client *elastic.Client, tokensMap map[string]Token) map[string]Token {
	bulk := client.Bulk().Index(indexName)
	for id, token := range tokensMap {
		item := elastic.NewBulkIndexRequest().Index(indexName).Id(id).OpType("index").Doc(token)
		bulk.Add(item)
	}
	response, err := bulk.Do(context.Background())
	if err != nil {
		log.Panicf("indexTokens() ERROR: %v\r\n", err)
	}
	if response.Errors {
		log.Panicf("indexTokens() ERROR: %v\r\n", response.Items)
	}
	_, _ = client.Flush(indexName).WaitIfOngoing(true).Do(context.Background())
	return tokensMap
}

//---- RetrieveCache subroutines --------------------------------------
func createFetchQuery(query RetrieveRequest) *elastic.BoolQuery {
	boolQuery := elastic.NewBoolQuery()
	userQuery := elastic.NewMatchQuery("user_id", query.UserId)
	if query.Type != BOTH {
		keywordQuery := elastic.NewMatchQuery("is_keyword", query.Type == KEYWORD)
		boolQuery.Must(userQuery, keywordQuery)
	} else {
		boolQuery.Must(userQuery)
	}
	hours := time.Duration(-query.Recency) * time.Hour
	from := GetCurrentLocalTime().Add(hours)
	boolQuery.Filter(elastic.NewRangeQuery("last_update").Gte(from))
	return boolQuery
}
func toCsv(searchResult *elastic.SearchResult, delimiter string) string {
	csv := ""
	//var tokens []string

	//for _, hit := range searchResult.Hits.Hits {
	//	var token Token
	//err := json.Unmarshal(hit.Source, &token)
	//if err != nil {
	//	log.Printf("toCsv() ERROR: %v\r\n", err)
	//}
	//tokens = append(tokens, token.Token)
	//}
	//if len(tokens) > 0 {
	//	csv = strings.Join(tokens, delimiter)
	//}
	return csv
}
