package route_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/ramadani/go-api-skeleton/app/user/route"
)

type UserRouteTestSuite struct {
	suite.Suite
	rr       *httptest.ResponseRecorder
	handlers *route.Handler
}

func (suite *UserRouteTestSuite) SetupTest() {
	suite.rr = httptest.NewRecorder()
	suite.handlers = route.NewHandler()
}

func (suite *UserRouteTestSuite) TestIndexRoute() {
	req, err := http.NewRequest(http.MethodGet, "/users", nil)
	suite.Nil(err)

	handler := http.HandlerFunc(suite.handlers.Index())
	handler.ServeHTTP(suite.rr, req)
	// Check the status code is what we expect.
	suite.Equal(http.StatusOK, suite.rr.Code)
	// Check the response body is what we expect.
	exceptedBody := `{"name":"Ramadani","email":"email.ramadani@gmail.com"}`
	suite.Equal(exceptedBody, suite.rr.Body.String())
}

func TestUserRouteTestSuite(t *testing.T) {
	suite.Run(t, new(UserRouteTestSuite))
}
