package create

import (
	"github.com/wagnerww/go-clean-arch/domain/customer"
)

type Create struct {
	repoSQL   customer.CustomerRepositoryInterface
	repoCache customer.CustomerRepositoryCacheInterface
}

func NewCreate(
	rSQL customer.CustomerRepositoryInterface,
	rCache customer.CustomerRepositoryCacheInterface) *Create {
	return &Create{
		repoSQL:   rSQL,
		repoCache: rCache,
	}
}

func (c *Create) CreateCustomer(input CustomerInputDto) (string, error) {
	cc, err := customer.NewCustomer("123", input.Name, input.PhoneNumber, input.Email)
	cc.Activate()

	if err != nil {
		return cc.ID, err
	}

	err = c.repoSQL.Create(cc)

	if err != nil {
		return cc.ID, err
	}

	c.repoCache.Create(cc.ID, cc)

	return cc.ID, nil

}
