#!/usr/bin/env bash
rm go-crypto-average-service 2> /dev/null
dep ensure
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-crypto-average-service .
docker build . -t go-crypto-average-service