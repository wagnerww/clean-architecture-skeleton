package create

import (
	"testing"

	"github.com/stretchr/testify/assert"
	customerGorm "github.com/wagnerww/go-clean-arch/infra/persistence/gorm/customer"
	customerRedis "github.com/wagnerww/go-clean-arch/infra/persistence/redis/customer"
)

func TestNewCustomer(t *testing.T) {
	repoSQL := customerGorm.NewCustomerRepository()
	repoCache := customerRedis.NewCustomerRepositoryCache()

	c := NewCreate(repoSQL, repoCache)

	input := CustomerInputDto{
		Email:       "email.mock@test.com.br",
		Name:        "customer name",
		PhoneNumber: 55692384,
	}
	_, err := c.CreateCustomer(input)
	assert.Nil(t, err)
}
