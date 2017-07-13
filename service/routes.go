package service

import (
	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc AppHandler
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/currencies/{symbol}",
		CurrencyHandler,
	}}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
		Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}
