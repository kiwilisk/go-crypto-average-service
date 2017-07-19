package service

import (
	"github.com/gorilla/mux"
	"net/http"
)

const symbolRequestParameter = "symbol"

type CurrencyHandler struct {
	floatingAverageService FloatingAverageService
}

func NewCurrencyHandler(floatingAverageService FloatingAverageService) *CurrencyHandler {
	return &CurrencyHandler{floatingAverageService}
}

func (handler CurrencyHandler) Handle(writer http.ResponseWriter, request *http.Request) error {
	vars := mux.Vars(request)
	symbol := vars[symbolRequestParameter]
	floatingAverageService := handler.floatingAverageService
	floatingAverage, err := floatingAverageService.Load(&symbol)
	if err != nil {
		return err
	}
	return WriteJsonResponse(writer, floatingAverage)
}
