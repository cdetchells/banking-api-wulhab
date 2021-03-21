package accounts

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"git.codesubmit.io/terem-technologies/banking-api-wulhab/internal/model"
)

type Repository interface {
	GetCustomerAccounts(id int) ([]*model.Account, error)
	GetCustomerAccount(id int) (*model.Account, error)
	CreateCustomerAccount(newAccount *model.NewAccount) (*model.Account, error)
}

func New() Repository {
	return &repository{}
}

type repository struct {
}

func (r *repository) GetCustomerAccounts(id int) ([]*model.Account, error) {
	accounts, _ := getAccounts()
	var customerAccounts []*model.Account
	for _, account := range accounts {
		if account.Customer == id {
			customerAccounts = append(customerAccounts, account)
		}
	}
	return customerAccounts, nil
}

func (r *repository) GetCustomerAccount(id int) (*model.Account, error) {
	accounts, _ := getAccounts()
	for _, account := range accounts {
		if account.ID == id {
			return account, nil
		}
	}
	return nil, nil
}

func (r *repository) CreateCustomerAccount(newAccount *model.NewAccount) (*model.Account, error) {
	accounts, _ := getAccounts()
	maxAccountId := 0
	for _, account := range accounts {
		if account.ID > maxAccountId {
			maxAccountId = account.ID
		}
	}
	return &model.Account{
		ID:       maxAccountId,
		Customer: newAccount.Customer,
		Type:     newAccount.Type,
		Balance:  newAccount.Balance,
	}, nil
}

func getAccounts() ([]*model.Account, error) {
	jsonFile, err := os.Open("./internal/data/accounts.json")
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var data []*model.Account
	json.Unmarshal(byteValue, &data)
	return data, nil
}
