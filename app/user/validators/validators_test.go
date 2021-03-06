package validators_test

import (
	"testing"

	"github.com/ramadani/go-api-skeleton/app/user/data"
	"github.com/ramadani/go-api-skeleton/app/user/validators"
	"github.com/stretchr/testify/assert"
)

func TestStoreValidator(t *testing.T) {
	user := data.UserInput{
		Name: "Great",
	}
	validator := validators.NewValidator()
	errs := validator.Validate(user)
	assert.Equal(t, 2, len(errs))
}
