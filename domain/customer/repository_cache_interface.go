package customer

import (
	"github.com/wagnerww/go-clean-arch/domain/shared/repository"
)

type CustomerRepositoryCacheInterface interface {
	repository.RepositoryCacheInterface[Customer]
}
