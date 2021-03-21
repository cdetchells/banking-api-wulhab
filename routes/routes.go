package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type route struct {
	pattern string
	handler func(http.ResponseWriter, *http.Request)
	method  string
}

func GetRoutes() []*route {
	return []*route{
		{
			pattern: "/customer/{id:[1-9]+}",
			handler: getCustomer,
			method:  "GET",
		},
		{
			pattern: "/customer/{id:[1-9]+}/accounts",
			handler: getCustomerAccounts,
			method:  "GET",
		},
	}
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	w.Write([]byte("Hello " + id))
}

func getCustomerAccounts(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	w.Write([]byte("Hello " + id))
}
