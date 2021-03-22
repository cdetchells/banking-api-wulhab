package datastore

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"git.codesubmit.io/terem-technologies/banking-api-wulhab/internal/model"
)

type DataStore interface {
	GetAccounts() ([]*model.Account, error)
	GetCustomers() ([]*model.Customer, error)
	GetTransactions() ([]*model.Transaction, error)
}

func New() DataStore {
	return &datastore{}
}

type datastore struct{}

func (d *datastore) GetAccounts() ([]*model.Account, error) {
	jsonFile, err := os.Open("./internal/data/accounts.json")
	if err != nil {
		return nil, err
	}
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	var data []*model.Account
	if err := json.Unmarshal(byteValue, &data); err != nil {
		return nil, err
	}
	return data, nil

}

func (d *datastore) GetCustomers() ([]*model.Customer, error) {
	jsonFile, err := os.Open("./internal/data/customers.json")
	if err != nil {
		return nil, err
	}
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	var customerData []*model.Customer
	if err := json.Unmarshal(byteValue, &customerData); err != nil {
		return nil, err
	}
	return customerData, nil
}

func (d *datastore) GetTransactions() ([]*model.Transaction, error) {
	jsonFile, err := os.Open("./internal/data/transactions.json")
	if err != nil {
		return nil, err
	}
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	var data []*model.Transaction
	if err := json.Unmarshal(byteValue, &data); err != nil {
		return nil, err
	}
	return data, nil
}
