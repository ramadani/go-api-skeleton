package route_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ramadani/go-api-skeleton/app/user/route"
	"github.com/stretchr/testify/assert"
)

func TestUserIndexHandler(t *testing.T) {
	assert := assert.New(t)
	req, err := http.NewRequest(http.MethodGet, "/users", nil)

	assert.NoError(err)

	rr := httptest.NewRecorder()
	h := route.NewHandler()

	handler := http.HandlerFunc(h.Index)
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	assert.Equal(http.StatusOK, rr.Code)

	// Check the response body is what we expect.
	exceptedBody := `{"name":"Ramadani","email":"email.ramadani@gmail.com"}`
	assert.Equal(exceptedBody, rr.Body.String())
}
