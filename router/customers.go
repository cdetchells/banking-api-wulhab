package router

import (
	"encoding/json"
	"net/http"
	"strconv"

	"git.codesubmit.io/terem-technologies/banking-api-wulhab/internal/repository/customers"
	"github.com/gorilla/mux"
)

func getCustomer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	// Do some basic validation on the request parameters
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid Customer Id"))
		return
	}
	customerRepo := customers.New()
	customer, err := customerRepo.GetCustomer(id)
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

func getCustomers(w http.ResponseWriter, r *http.Request) {
	customerRepo := customers.New()
	customer, _ := customerRepo.GetCustomers()
	c, _ := json.Marshal(customer)
	w.Write(c)
}
