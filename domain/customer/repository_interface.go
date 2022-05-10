package customer

import (
	"github.com/wagnerww/go-clean-arch/domain/shared/repository"
)

type CustomerRepositoryInterface interface {
	repository.RepositoryInterface[Customer]
}
