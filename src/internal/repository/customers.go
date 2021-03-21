package customers

type Repository interface {
	GetCustomers() ([]*Customer, error)
	GetCustomer(id int32) (Customer, error)
}

func New() Repository {
	return &repository{}
}

type repository struct {
}

func (r *repository) GetCustomers() ([]*Customer, error) {
	return
}

func (r *repository) GetCustomer() ([]*Customer, error) {

}
