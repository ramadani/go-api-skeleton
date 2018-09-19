package route_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/ramadani/go-api-skeleton/app/welcome/route"
)

type WelcomeRouteTestSuite struct {
	suite.Suite
	rr       *httptest.ResponseRecorder
	handlers *route.Handler
}

func (suite *WelcomeRouteTestSuite) SetupTest() {
	suite.rr = httptest.NewRecorder()
	suite.handlers = route.NewHandler()
}

func (suite *WelcomeRouteTestSuite) TestIndexRoute() {
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	suite.Nil(err)

	handler := http.HandlerFunc(suite.handlers.Index())
	handler.ServeHTTP(suite.rr, req)
	// Check the status code is what we expect.
	suite.Equal(http.StatusOK, suite.rr.Code)
	// Check the response body is what we expect.
	expectedBody := `{"name":"Go API Skeleton","description":"Go (Golang) API Skeleton for your great API","version":"v0.1.0"}`
	suite.Equal(expectedBody, suite.rr.Body.String())
}

func TestWelcomeRouteTestSuite(t *testing.T) {
	suite.Run(t, new(WelcomeRouteTestSuite))
}
