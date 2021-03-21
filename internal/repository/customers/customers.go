package customers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"git.codesubmit.io/terem-technologies/banking-api-wulhab/internal/model"
)

type Repository interface {
	GetCustomers() ([]*model.Customer, error)
	GetCustomer(id int) (*model.Customer, error)
}

func New() Repository {
	return &repository{}
}

type repository struct {
}

func (r *repository) GetCustomers() ([]*model.Customer, error) {
	customers, _ := getCustomers()
	return customers, nil
}

func (r *repository) GetCustomer(id int) (*model.Customer, error) {
	customers, _ := getCustomers()
	for _, customer := range customers {
		if customer.ID == id {
			return customer, nil
		}
	}
	return nil, nil
}

func getCustomers() ([]*model.Customer, error) {
	jsonFile, err := os.Open("./internal/data/customers.json")
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var customerData []*model.Customer
	json.Unmarshal(byteValue, &customerData)
	return customerData, nil
}
