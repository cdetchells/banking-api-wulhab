package router

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (rtr *router) GetCustomer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	// Do some basic validation on the request parameters
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid Customer Id"))
		return
	}
	customer, err := rtr.customers.GetCustomer(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if customer == nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Customer Not Found"))
		return
	}
	c, _ := json.Marshal(customer)
	w.Write(c)
}

func (rtr *router) GetCustomers(w http.ResponseWriter, r *http.Request) {
	customer, _ := rtr.customers.GetCustomers()
	c, _ := json.Marshal(customer)
	w.Write(c)
}
