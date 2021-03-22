package accounts

import (
	"git.codesubmit.io/terem-technologies/banking-api-wulhab/internal/datastore"
	"git.codesubmit.io/terem-technologies/banking-api-wulhab/internal/model"
)

type Repository interface {
	CreateCustomerAccount(newAccount *model.NewAccount) (*model.Account, error)
	GetAccount(id int) (*model.Account, error)
	GetCustomerAccount(customerId int, id int) (*model.Account, error)
	GetCustomerAccounts(id int) ([]*model.Account, error)
}

func New() Repository {
	return &repository{datastore: datastore.New()}
}

type repository struct {
	datastore datastore.DataStore
}

func (r *repository) GetCustomerAccounts(id int) ([]*model.Account, error) {
	accounts, err := r.datastore.GetAccounts()
	if err != nil {
		return nil, err
	}
	var customerAccounts []*model.Account
	for _, account := range accounts {
		if account.Customer == id {
			customerAccounts = append(customerAccounts, account)
		}
	}
	return customerAccounts, nil
}

func (r *repository) GetCustomerAccount(customerId int, id int) (*model.Account, error) {
	accounts, err := r.datastore.GetAccounts()
	if err != nil {
		return nil, err
	}
	for _, account := range accounts {
		if account.ID == id && account.Customer == customerId {
			return account, nil
		}
	}
	return nil, nil
}

// Then we mimick creating an account but all it does is get the next account id and pass the same data back to the user
func (r *repository) CreateCustomerAccount(newAccount *model.NewAccount) (*model.Account, error) {
	accounts, err := r.datastore.GetAccounts()
	if err != nil {
		return nil, err
	}
	maxAccountId := 0
	for _, account := range accounts {
		if account.ID > maxAccountId {
			maxAccountId = account.ID
		}
	}
	return &model.Account{
		ID:       maxAccountId + 1,
		Customer: newAccount.Customer,
		Type:     newAccount.Type,
		Balance:  newAccount.Balance,
	}, nil
}

func (r *repository) GetAccount(id int) (*model.Account, error) {
	accounts, err := r.datastore.GetAccounts()
	if err != nil {
		return nil, err
	}
	for _, account := range accounts {
		if account.ID == id {
			return account, nil
		}
	}
	return nil, nil
}
