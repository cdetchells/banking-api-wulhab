package router

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"git.codesubmit.io/terem-technologies/banking-api-wulhab/internal/model"
	"git.codesubmit.io/terem-technologies/banking-api-wulhab/internal/repository/transactions"
	"github.com/gorilla/mux"
)

func createAccountTransfer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	// Do some basic validation on the request parameters
	customerId, err := strconv.Atoi(params["customerid"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid Customer Id"))
		return
	}
	accountId, err := strconv.Atoi(params["accountid"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid Account Id"))
		return
	}

	var newTransaction *model.NewTransaction
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid Request Body"))
		return
	}
	err = json.Unmarshal(body, &newTransaction)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid Request Body"))
		return
	}
	repo := transactions.New()
	data, err := repo.CreateAccountTransaction(customerId, accountId, newTransaction)
	if err != nil {
		if err.Error() == "To Account Not Found" || err.Error() == "Account Not Found" || err.Error() == "Cannot Transfer To Same Account" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	d, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(d)
}

func getAccountTransactions(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	// Do some validation on the request parameters
	customerId, err := strconv.Atoi(params["customerid"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid Customer Id"))
		return
	}
	accountId, err := strconv.Atoi(params["accountid"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid Account Id"))
		return
	}
	repo := transactions.New()
	data, err := repo.GetAccountTransactions(customerId, accountId)
	if err != nil {
		if err.Error() == "Account Not Found" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	d, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(d)
}
