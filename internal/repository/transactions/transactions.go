package transactions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"git.codesubmit.io/terem-technologies/banking-api-wulhab/internal/model"
)

type Repository interface {
	GetAccountTransactions(id int) ([]*model.Transaction, error)
	CreateAccountTransaction(newTransaction *model.NewTransaction) (*model.Transaction, error)
}

func New() Repository {
	return &repository{}
}

type repository struct {
}

func (r *repository) GetAccountTransactions(id int) ([]*model.Transaction, error) {
	transactions, _ := getTransactions()
	var accountTrans []*model.Transaction
	for _, transaction := range transactions {
		fmt.Println(id)
		fmt.Println(transaction.FromAccount)
		if transaction.FromAccount == id || transaction.ToAccount == id {
			accountTrans = append(accountTrans, transaction)
		}
	}
	return accountTrans, nil
}

func (r *repository) CreateAccountTransaction(newTransaction *model.NewTransaction) (*model.Transaction, error) {
	transactions, _ := getTransactions()
	maxTransId := 0
	for _, transaction := range transactions {
		if transaction.ID > maxTransId {
			maxTransId = transaction.ID
		}
	}
	return &model.Transaction{
		ID:          maxTransId + 1,
		CreatedAt:   newTransaction.CreatedAt,
		FromAccount: newTransaction.FromAccount,
		ToAccount:   newTransaction.ToAccount,
		Amount:      newTransaction.Amount,
	}, nil
}

func getTransactions() ([]*model.Transaction, error) {
	jsonFile, err := os.Open("./internal/data/transactions.json")
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var data []*model.Transaction
	json.Unmarshal(byteValue, &data)
	return data, nil
}
