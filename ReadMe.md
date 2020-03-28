# Cache Service

Simple cache management service written in Golang.

## Requirement
Install Elasticsearch +7.6 <p/>
[Insatlling Elasticsearch](https://www.elastic.co/guide/en/elasticsearch/reference/current/install-elasticsearch.html)<p/>
[Install Elasticsearch with Docker](https://www.elastic.co/guide/en/elasticsearch/reference/current/docker.html)<p/>
[Install ELKStack using docker-compose](https://github.com/deviantony/docker-elk)
## Running
Run following commands:
````
git clone https://github.com/vhesami/cache-service
cd ~/cache-service
sudo docker build -t cache-service .
sudo docker run -t -i -e PORT='8080' -e ES_URL='http://msvc.ir:9200' -p 8080:8080 --name simple_cache_service --rm cache-service
````
**"PORT"** is the web service listening port number and **"ES_URL"** is the Elasticsearch service url.
## APIs
1.**StoreCache**: This api provides stores user's queries in the cache.
````
curl --header "Content-Type: application/json" \
     --request POST \
     --data '{"user_id":"user6985","text":"Go how to program!"}' \
    http://localhost:8080/cache/api/v1/store
````
**user_id** is the specific user's id and **text** is the query of user than should be store.

Result:
````
{"success":true,"stored_token_count":5}
````

2.**RetrieveCache**: This api provides frequent tokens of specific user in last defined period in csv format.
````
curl --header "Content-Type: application/json" \
     --request POST \
     --data '{"user_id":"user6985","hours":1,"size":10, "type":1, "delimiter":","}' \
    http://localhost:8080/cache/api/v1/retrieve
````
**user_id** is the specific user's id.

**hours** is the number of hours in time periods. (default is 1)

**size** is the number of returning results. (default is 10)

**type** is the type of returning tokens. 1 => STANDARD, 2=> KEYWORD and 3=> BOTH. (default is STANDARD)

**delimiter** is the delimiter sign in returning results. (default is ',')

Result:
````
to,program,Go,how
````
## Architecture
This service designed by MRU Cache idea and includes two APIs:
1. Store In Cache
2. Retrieve From Cache

In the first API, text of query tokenized by two strategies; [Standard](https://lucene.apache.org/solr/guide/6_6/tokenizers.html#Tokenizers-StandardTokenizer) and [Keyword](https://lucene.apache.org/solr/guide/6_6/tokenizers.html#Tokenizers-KeywordTokenizer).
In the next step user's tokens gets unique identifier and store in the ElasticSearch using [bulk api](https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-bulk.html).
In the next appear of each token the "occurs" and "last_update" fields of that token will be updated.

In the "Retrieve" API, an ElasticSearch query will be run with "user_id" and another specific parameters.
Finally the results sort by "occurs" field descending. In this scenario the Most Recently Used tokens become as result.

In this architecture [ElasticSearch](https://www.elastic.co)(powered by [Lucene](https://lucene.apache.org/)) has key role to acquire acceptable performance.