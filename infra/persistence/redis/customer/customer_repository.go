package customer

import (
	"fmt"

	"github.com/wagnerww/go-clean-arch/domain/customer"
)

var KEY string = "CUSTOMER_"

type CustomerRepositoryCache struct {
}

func NewCustomerRepositoryCache() *CustomerRepositoryCache {
	return &CustomerRepositoryCache{}
}

func (c *CustomerRepositoryCache) Create(keyId string, entity customer.Customer) error {
	fmt.Println("chegou aqui Create CACHE" + KEY + keyId)
	return nil
}

func (c *CustomerRepositoryCache) FindById(keyId string) (customer.Customer, error) {
	fmt.Println("chegou aqui find CACHE" + KEY + keyId)
	return customer.Customer{}, nil
}

func (c *CustomerRepositoryCache) DeleteById(keyId string) error {
	fmt.Println("chegou aqui Delete" + KEY + keyId)
	return nil
}
