package tests

import (
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
}

func TestUserRouteTestSuite(t *testing.T) {
	suite.Run(t, new(UserRouteTestSuite))
}
