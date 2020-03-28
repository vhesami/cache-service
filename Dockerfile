FROM golang

ADD . /go/src/github.com/vhesami/cache-service
WORKDIR . /go/src/github.com/vhesami/cache-service

RUN go get github.com/gorilla/mux
RUN go get github.com/olivere/elastic

RUN go install github.com/vhesami/cache-service
ENTRYPOINT /go/bin/cache-service
