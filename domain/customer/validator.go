package customer

import (
	"errors"
)

func Validate(entity Customer) error {
	if entity.Name == "" {
		return errors.New("Name is required")
	}
	return nil
}
