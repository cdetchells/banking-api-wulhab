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
			handler: customer,
			method:  "GET",
		},
	}
}

func customer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]
	w.Write([]byte("Hello " + name))
}
