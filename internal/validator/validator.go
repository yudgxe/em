package validator

import "github.com/go-playground/validator/v10"

var val = validator.New()

func Validate(in any) error {
	return val.Struct(in)
}
