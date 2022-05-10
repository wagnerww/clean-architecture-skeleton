package sql

import (
	"fmt"

	"github.com/wagnerww/go-clean-arch/domain/customer"
)

type CustomerRepository struct {
}

func NewCustomerRepository() *CustomerRepository {
	return &CustomerRepository{}
}

func (c *CustomerRepository) Create(entity customer.Customer) error {
	customer := Customer{
		ID:          entity.ID,
		Name:        entity.Name,
		PhoneNumber: entity.PhoneNumber,
		Email:       entity.Email,
		Active:      entity.IsActive(),
	}

	fmt.Println("chegou aqui Create", customer)
	return nil
}

func (c *CustomerRepository) Update(entity customer.Customer) error {
	fmt.Println("chegou aqui Update")
	return nil
}

func (c *CustomerRepository) FindById(id string) (customer.Customer, error) {

	fmt.Println("chegou aqui findbyid")
	return customer.Customer{}, nil
}

func (c *CustomerRepository) FindAll() ([]customer.Customer, error) {
	fmt.Println("chegou aqui findAll")
	return []customer.Customer{}, nil
}
