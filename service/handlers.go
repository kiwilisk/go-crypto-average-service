package service

import (
	"net/http"
)

type AppHandler func(http.ResponseWriter, *http.Request) error

func (handlerFunction AppHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	if err := handlerFunction(responseWriter, request); err != nil {
		http.Error(responseWriter, err.Error(), 500)
	}
}
