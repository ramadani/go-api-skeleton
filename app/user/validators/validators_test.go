package validators_test

import (
	"testing"

	"github.com/ramadani/go-api-skeleton/app/user/validators"
	"github.com/stretchr/testify/assert"
)

func TestStoreValidator(t *testing.T) {
	user := validators.UserInput{
		Name: "Great",
	}
	validator := validators.NewValidator()
	errs := validator.Store(user)
	assert.Equal(t, 2, len(errs))
}
