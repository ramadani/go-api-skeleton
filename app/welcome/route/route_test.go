package route_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ramadani/go-api-skeleton/app/welcome/route"
	"github.com/stretchr/testify/assert"
)

func TestWelcomeIndexHandler(t *testing.T) {
	assert := assert.New(t)
	req, err := http.NewRequest(http.MethodGet, "/", nil)

	assert.NoError(err)

	rr := httptest.NewRecorder()
	h := route.NewHandler()

	handler := http.HandlerFunc(h.Index)
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	statusCode := rr.Code
	assert.Equal(http.StatusOK, statusCode)

	// Check the response body is what we expect.
	expectedBody := `{"name":"Go API Skeleton","description":"Go (Golang) API Skeleton for your great API","version":"v0.1.0"}`
	actualBody := rr.Body.String()
	assert.Equal(expectedBody, actualBody)
}
