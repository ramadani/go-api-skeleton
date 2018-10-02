package validator

import (
	"github.com/ramadani/go-api-skeleton/commons/str"
	validator "gopkg.in/go-playground/validator.v9"
)

// Errors transformer
func Errors(errs validator.ValidationErrors) map[string]interface{} {
	var errors = make(map[string]interface{}, len(errs))
	for _, err := range errs {
		var fieldErr = make(map[string]string)
		fieldErr["rule"] = err.Tag()
		if err.Param() != "" {
			fieldErr["param"] = err.Param()
		}
		errors[str.ToSnakeCase(err.Field())] = fieldErr
	}
	return errors
}
