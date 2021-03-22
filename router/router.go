package router

import (
	"net/http"

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

func New() http.Handler {
	gRtr := mux.NewRouter()

	router := &router{customers: customers.New(), accounts: accounts.New(), transactions: transactions.New()}
	routes := router.GetRoutes()
	for _, route := range routes {
		gRtr.HandleFunc(route.pattern, route.handler).Methods(route.method)
	}
	return gRtr
}

type router struct {
	customers    customers.Repository
	accounts     accounts.Repository
	transactions transactions.Repository
}

func (r *router) GetRoutes() []*route {
	return []*route{
		{
			pattern: "/customers",
			handler: r.GetCustomers,
			method:  "GET",
		},
		{
			pattern: "/customers/{id:[1-9]+}",
			handler: r.GetCustomer,
			method:  "GET",
		},
		{
			pattern: "/customers/{id:[1-9]+}/accounts",
			handler: r.GetCustomerAccounts,
			method:  "GET",
		},
		{
			pattern: "/customers/{customerid:[1-9]+}/accounts/{accountid:[1-9]+}",
			handler: r.GetCustomerAccount,
			method:  "GET",
		},
		{
			pattern: "/customers/{id:[1-9]+}/accounts",
			handler: r.CreateCustomerAccount,
			method:  "POST",
		},
		{
			pattern: "/customers/{customerid:[1-9]+}/accounts/{accountid:[1-9]+}/transactions",
			handler: r.CreateAccountTransfer,
			method:  "POST",
		},
		{
			pattern: "/customers/{customerid:[1-9]+}/accounts/{accountid:[1-9]+}/transactions",
			handler: r.GetAccountTransactions,
			method:  "GET",
		},
	}
}
