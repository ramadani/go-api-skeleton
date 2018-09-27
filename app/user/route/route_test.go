package route_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/ramadani/go-api-skeleton/app/user/data"
	"github.com/ramadani/go-api-skeleton/app/user/route"
	"github.com/ramadani/go-api-skeleton/app/user/usecase/mocks"
	hl "github.com/ramadani/go-api-skeleton/commons/handler"
)

type UserRouteTestSuite struct {
	suite.Suite
	rr       *httptest.ResponseRecorder
	handlers *route.Handler
}

func (suite *UserRouteTestSuite) SetupTest() {
	suite.rr = httptest.NewRecorder()
}

func (suite *UserRouteTestSuite) TestIndexRoute() {
	var page, limit, total uint
	page = 1
	limit = 10
	total = 10
	users := make([]data.User, limit)
	uri := fmt.Sprintf("/users?page=%d", page)

	req, err := http.NewRequest(http.MethodGet, uri, nil)
	suite.Nil(err)

	ucase := new(mocks.Usecase)
	suite.handlers = route.NewHandler(ucase)

	defer ucase.AssertExpectations(suite.T())

	for i := uint(0); i < limit; i++ {
		users[i] = data.User{
			ID:    i + 1,
			Name:  fmt.Sprintf("FooBar %d", i),
			Email: fmt.Sprintf("foo%d@example.com", i),
		}
	}

	userPaginate := data.UserPaginate{
		Data:    users,
		Total:   total,
		PerPage: limit,
		Page:    1,
		Pages:   1,
	}

	ucase.On("Paginate", page, limit).Return(userPaginate, nil).Once()

	handler := http.HandlerFunc(suite.handlers.Index())
	handler.ServeHTTP(suite.rr, req)
	exceptedBody, _ := json.Marshal(hl.ResponseData{Data: userPaginate})
	suite.Equal(string(exceptedBody), suite.rr.Body.String())
	suite.Equal(http.StatusOK, suite.rr.Code)
}

func (suite *UserRouteTestSuite) TestIndexRouteOnFailed() {
	var page, limit uint
	page = 1
	limit = 10

	req, err := http.NewRequest(http.MethodGet, "/users", nil)
	suite.Nil(err)

	ucase := new(mocks.Usecase)
	suite.handlers = route.NewHandler(ucase)

	defer ucase.AssertExpectations(suite.T())

	ucase.On("Paginate", page, limit).Return(data.UserPaginate{}, fmt.Errorf("internal server error")).Once()
	resErr := hl.ResponseError{Message: "internal server error"}

	handler := http.HandlerFunc(suite.handlers.Index())
	handler.ServeHTTP(suite.rr, req)
	exceptedBody, _ := json.Marshal(hl.ResponseData{Data: resErr})
	suite.Equal(string(exceptedBody), suite.rr.Body.String())
	suite.Equal(http.StatusInternalServerError, suite.rr.Code)
}

func TestUserRouteTestSuite(t *testing.T) {
	suite.Run(t, new(UserRouteTestSuite))
}
