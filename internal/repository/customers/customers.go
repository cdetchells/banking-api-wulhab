package customers

import (
	"git.codesubmit.io/terem-technologies/banking-api-wulhab/internal/datastore"
	"git.codesubmit.io/terem-technologies/banking-api-wulhab/internal/model"
)

type Repository interface {
	GetCustomers() ([]*model.Customer, error)
	GetCustomer(id int) (*model.Customer, error)
}

func New() Repository {
	return &repository{datastore: datastore.New()}
}

type repository struct {
	datastore datastore.DataStore
}

func (r *repository) GetCustomers() ([]*model.Customer, error) {
	customers, err := r.datastore.GetCustomers()
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (r *repository) GetCustomer(id int) (*model.Customer, error) {
	customers, err := r.datastore.GetCustomers()
	if err != nil {
		return nil, err
	}
	for _, customer := range customers {
		if customer.ID == id {
			return customer, nil
		}
	}
	return nil, nil
}
