package route_test

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"
)

type UserRouteTestSuite struct {
	suite.Suite
	rr *httptest.ResponseRecorder
}

func (suite *UserRouteTestSuite) SetupTest() {
	suite.rr = httptest.NewRecorder()
}

func TestUserRouteTestSuite(t *testing.T) {
	suite.Run(t, new(UserRouteTestSuite))
}
