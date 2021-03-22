package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

type route struct {
	pattern string
	handler func(http.ResponseWriter, *http.Request)
	method  string
}

func New() http.Handler {
	rtr := mux.NewRouter()
	routes := getRoutes()
	for _, route := range routes {
		rtr.HandleFunc(route.pattern, route.handler).Methods(route.method)
	}
	return rtr
}

func getRoutes() []*route {
	return []*route{
		{
			pattern: "/customers",
			handler: getCustomers,
			method:  "GET",
		},
		{
			pattern: "/customers/{id:[1-9]+}",
			handler: getCustomer,
			method:  "GET",
		},
		{
			pattern: "/customers/{id:[1-9]+}/accounts",
			handler: getCustomerAccounts,
			method:  "GET",
		},
		{
			pattern: "/customers/{customerid:[1-9]+}/accounts/{accountid:[1-9]+}",
			handler: getCustomerAccount,
			method:  "GET",
		},
		{
			pattern: "/customers/{id:[1-9]+}/accounts",
			handler: createCustomerAccount,
			method:  "POST",
		},
		{
			pattern: "/customers/{customerid:[1-9]+}/accounts/{accountid:[1-9]+}/transactions",
			handler: createAccountTransfer,
			method:  "POST",
		},
		{
			pattern: "/customers/{customerid:[1-9]+}/accounts/{accountid:[1-9]+}/transactions",
			handler: getAccountTransactions,
			method:  "GET",
		},
	}
}
