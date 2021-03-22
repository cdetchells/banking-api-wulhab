package transactions

import (
	"errors"

	"git.codesubmit.io/terem-technologies/banking-api-wulhab/internal/datastore"
	"git.codesubmit.io/terem-technologies/banking-api-wulhab/internal/model"
	"git.codesubmit.io/terem-technologies/banking-api-wulhab/internal/repository/accounts"
)

type Repository interface {
	GetAccountTransactions(customerId int, accountId int) ([]*model.Transaction, error)
	CreateAccountTransaction(customerId int, fromAccountId int, newTransaction *model.NewTransaction) (*model.Transaction, error)
}

func New() Repository {
	return &repository{datastore: datastore.New(), accounts: accounts.New()}
}

type repository struct {
	datastore datastore.DataStore
	accounts  accounts.Repository
}

func (r *repository) GetAccountTransactions(customerId int, accountId int) ([]*model.Transaction, error) {
	account, err := r.accounts.GetCustomerAccount(customerId, accountId)
	if err != nil {
		return nil, err
	}
	if account == nil {
		return nil, errors.New("Account Not Found")
	}

	transactions, err := r.datastore.GetTransactions()
	if err != nil {
		return nil, err
	}
	var accountTrans []*model.Transaction
	for _, transaction := range transactions {
		if transaction.FromAccount == accountId || transaction.ToAccount == accountId {
			// If it's From this account then it should be a debit
			if transaction.FromAccount == accountId {
				transaction.Amount = 0 - transaction.Amount
			}
			accountTrans = append(accountTrans, transaction)
		}
	}
	return accountTrans, nil
}

func (r *repository) CreateAccountTransaction(customerId int, fromAccountId int, newTransaction *model.NewTransaction) (*model.Transaction, error) {
	// Check that the from account exists for the customer
	fromAccount, err := r.accounts.GetCustomerAccount(customerId, fromAccountId)
	if err != nil {
		return nil, err
	}
	if fromAccount == nil {
		return nil, errors.New("Account Not Found")
	}

	// Check that the ToAccount exists
	toAccount, err := r.accounts.GetAccount(newTransaction.ToAccount)
	if err != nil {
		return nil, err
	}
	if toAccount == nil {
		return nil, errors.New("To Account Not Found")
	}

	// From and To accounts must be different
	if newTransaction.ToAccount == fromAccountId {
		return nil, errors.New("Cannot Transfer To Same Account")
	}

	// Then we mimick creating a transaction but all it does is get the next transaction id and pass the data back to the user
	transactions, err := r.datastore.GetTransactions()
	if err != nil {
		return nil, err
	}
	maxTransId := 0
	for _, transaction := range transactions {
		if transaction.ID > maxTransId {
			maxTransId = transaction.ID
		}
	}
	return &model.Transaction{
		ID:          maxTransId + 1,
		CreatedAt:   newTransaction.CreatedAt,
		FromAccount: fromAccountId,
		ToAccount:   newTransaction.ToAccount,
		Amount:      newTransaction.Amount,
	}, nil
}

//go:generate mockgen -source=transactions.go -package=transactions -destination=./mocks/mock_transactions.go -package=mocks
