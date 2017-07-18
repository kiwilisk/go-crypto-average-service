package service

import (
	"log"
	"net/http"
	"time"
)

type AppHandler func(http.ResponseWriter, *http.Request) error

func (handlerFunction AppHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	if err := handlerFunction(responseWriter, request); err != nil {
		http.Error(responseWriter, err.Error(), 500)
	}
}

func LogRequest(inner AppHandler, name string) AppHandler {
	return func(responseWriter http.ResponseWriter, request *http.Request) error {
		start := time.Now()
		err := inner(responseWriter, request)
		log.Printf(
			"%s\t%s\t%s\t%s\tsuccessful:%t",
			request.Method,
			request.RequestURI,
			name,
			time.Since(start),
			err == nil,
		)
		return err
	}
}
