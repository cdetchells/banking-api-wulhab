package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"git.codesubmit.io/terem-technologies/banking-api-wulhab/internal/model"
	"git.codesubmit.io/terem-technologies/banking-api-wulhab/internal/repository/accounts"
	"git.codesubmit.io/terem-technologies/banking-api-wulhab/internal/repository/customers"
	"git.codesubmit.io/terem-technologies/banking-api-wulhab/internal/repository/transactions"
	"github.com/gorilla/mux"
)

type route struct {
	pattern string
	handler func(http.ResponseWriter, *http.Request)
	method  string
}

func main() {
	rtr := mux.NewRouter()
	routes := []*route{
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
			pattern: "/customers/{id:[1-9]+}/accounts",
			handler: createCustomerAccount,
			method:  "POST",
		},
		{
			pattern: "/customers/{customerid:[1-9]+}/accounts/{accountid:[1-9]+}",
			handler: createAccountTransfer,
			method:  "POST",
		},
		{
			pattern: "/customers/{customerid:[1-9]+}/accounts/{accountid:[1-9]+}/transactions",
			handler: getAccountTransactions,
			method:  "GET",
		},
	}
	for _, route := range routes {
		rtr.HandleFunc(route.pattern, route.handler).Methods(route.method)
	}

	http.Handle("/", rtr)

	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	customerRepo := customers.New()
	customer, _ := customerRepo.GetCustomer(id)
	c, _ := json.Marshal(customer)
	w.Write(c)
}

func getCustomers(w http.ResponseWriter, r *http.Request) {
	customerRepo := customers.New()
	customer, _ := customerRepo.GetCustomers()
	c, _ := json.Marshal(customer)
	w.Write(c)
}

func getCustomerAccounts(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	repo := accounts.New()
	data, _ := repo.GetCustomerAccounts(id)
	d, _ := json.Marshal(data)
	w.Write(d)
}

func createCustomerAccount(w http.ResponseWriter, r *http.Request) {
	var newAccount *model.NewAccount
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &newAccount)
	repo := accounts.New()
	data, _ := repo.CreateCustomerAccount(newAccount)
	d, _ := json.Marshal(data)
	w.Write(d)
}

func createAccountTransfer(w http.ResponseWriter, r *http.Request) {
	var newTransaction *model.NewTransaction
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &newTransaction)
	repo := transactions.New()
	data, _ := repo.CreateAccountTransaction(newTransaction)
	d, _ := json.Marshal(data)
	w.Write(d)
}

func getAccountTransactions(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["accountid"])
	repo := transactions.New()
	data, _ := repo.GetAccountTransactions(id)
	d, _ := json.Marshal(data)
	w.Write(d)
}
