package validator

import (
	v "github.com/go-playground/validator/v10"
)

type structValidator struct {
	validate *v.Validate
}

func NewStructValidator() *structValidator {
	return &structValidator{
		validate: v.New(v.WithRequiredStructEnabled()),
	}
}

// Validator needs to implement the Validate method
func (v *structValidator) Validate(out any) error {
	return v.validate.Struct(out)
}
