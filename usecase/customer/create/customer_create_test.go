package create

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wagnerww/go-clean-arch/infra/repository/customer/cache"
	"github.com/wagnerww/go-clean-arch/infra/repository/customer/sql"
)

func TestNewCustomer(t *testing.T) {
	repoSQL := sql.NewCustomerRepository()
	repoCache := cache.NewCustomerRepositoryCache()

	c := NewCreate(repoSQL, repoCache)

	input := CustomerInputDto{
		Email:       "email.mock@test.com.br",
		Name:        "customer name",
		PhoneNumber: 55692384,
	}
	_, err := c.CreateCustomer(input)
	assert.Nil(t, err)
}
