package repository

type RepositoryCacheInterface[T any] interface {
	Create(keyId string, entity T) error
	FindById(keyId string) (T, error)
	DeleteById(keyId string) error
}
