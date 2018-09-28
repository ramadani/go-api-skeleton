package validators

import (
	helper "github.com/ramadani/go-api-skeleton/commons/validator"
	validator "gopkg.in/go-playground/validator.v9"
)

// UserInput for validate the input
type UserInput struct {
	Name     string `validate:"required"`
	Email    string `validate:"required,lt=10"`
	Password string `validate:"required"`
}

// Validator for user input
type Validator struct {
	v *validator.Validate
}

// Store user validator
func (vdt *Validator) Store(input UserInput) map[string]interface{} {
	var errors = make(map[string]interface{})
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
