package validators

import (
	helper "github.com/ramadani/go-api-skeleton/commons/validator"
	validator "gopkg.in/go-playground/validator.v9"
)

// Validator for user input
type Validator struct {
	v *validator.Validate
}

// Validate user validator
func (vdt *Validator) Validate(input interface{}) (errors map[string]interface{}) {
	err := vdt.v.Struct(input)
	if err != nil {
		errors = helper.Errors(err.(validator.ValidationErrors))
	}

	return errors
}

// NewValidator for user
func NewValidator() *Validator {
	return &Validator{validator.New()}
}
