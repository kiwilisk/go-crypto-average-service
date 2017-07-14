package service

import (
	"github.com/gorilla/mux"
	"net/http"
)

const symbolRequestParameter = "symbol"

func CurrencyHandler(writer http.ResponseWriter, request *http.Request) error {
	vars := mux.Vars(request)
	symbol := vars[symbolRequestParameter]
	floatingAverage, err := LoadFloatingAverage(&symbol)
	if err != nil {
		return err
	}
	return WriteJsonResponse(writer, floatingAverage)
}
