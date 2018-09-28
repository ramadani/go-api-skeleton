package validators

import (
	"github.com/ramadani/go-api-skeleton/app/user/data"
	helper "github.com/ramadani/go-api-skeleton/commons/validator"
	validator "gopkg.in/go-playground/validator.v9"
)

// Validator for user input
type Validator struct {
	v *validator.Validate
}

// Store user validator
func (vdt *Validator) Store(input data.UserInput) (errors map[string]interface{}) {
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
