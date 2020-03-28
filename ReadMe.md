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
## Services
1. **StoreCache**: This api provides stores user's queries in the cache.
````

````