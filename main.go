package main

import (
	"github.com/kiwilisk/go-crypto-average-service/service"
	"log"
	"net/http"
)

func main() {
	router := service.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
