package validator

type ValidatorInterface[T any] interface {
	Validate(entity T) error
}
