package router

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"git.codesubmit.io/terem-technologies/banking-api-wulhab/internal/model"
	"git.codesubmit.io/terem-technologies/banking-api-wulhab/internal/repository/accounts"
	"github.com/gorilla/mux"
)

func getCustomerAccounts(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	// Do some basic validation on the request parameters
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid Account Id"))
		return
	}
	repo := accounts.New()
	data, _ := repo.GetCustomerAccounts(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if data == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	d, _ := json.Marshal(data)
	w.Write(d)
}

func getCustomerAccount(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	// Do some basic validation on the request parameters
	customerid, err := strconv.Atoi(params["customerid"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid Customer Id"))
		return
	}
	accountid, err := strconv.Atoi(params["accountid"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid Account Id"))
		return
	}
	repo := accounts.New()
	data, _ := repo.GetCustomerAccount(customerid, accountid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if data == nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Account Not Found"))
		return
	}
	d, _ := json.Marshal(data)
	w.Write(d)
}

func createCustomerAccount(w http.ResponseWriter, r *http.Request) {
	var newAccount *model.NewAccount
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(body, &newAccount)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	repo := accounts.New()
	data, err := repo.CreateCustomerAccount(newAccount)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if data == nil {
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
