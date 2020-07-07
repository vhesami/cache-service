FROM golang:1.11.4-alpine3.8 AS build
ENV GO111MODULE=on \
    GOOS=linux \
    GOARCH=amd64
RUN apk add --update --no-cache \
      build-base \
      git \
      ca-certificates \
    && \
    mkdir -p /src
COPY go.sum go.mod /src/
WORKDIR /src
RUN go mod download
COPY . /src
RUN go build -ldflags="-w -s" -o cache-service


FROM alpine:3.8
ENV TZ=UTC
RUN apk add --update --no-cache \
      tzdata \
      ca-certificates \
    && \
    cp --remove-destination /usr/share/zoneinfo/${TZ} /etc/localtime && \
    echo "${TZ}" > /etc/timezone
WORKDIR /app
EXPOSE 8080
COPY --from=build /src/cache-service /app/
CMD ["/app/cache-service"]