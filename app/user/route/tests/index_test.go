package tests

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/ramadani/go-api-skeleton/app/user/data"
	"github.com/ramadani/go-api-skeleton/app/user/route"
	"github.com/ramadani/go-api-skeleton/app/user/usecase/mocks"
	"github.com/ramadani/go-api-skeleton/commons/http/res"
)

func (suite *UserRouteTestSuite) TestIndexRoute() {
	ucase := new(mocks.Usecase)
	handlers := route.NewHandler(ucase)

	defer ucase.AssertExpectations(suite.T())

	var page, limit, total uint
	page = 1
	limit = 10
	total = 10
	users := make([]data.User, limit)
	uri := fmt.Sprintf("/users?page=%d", page)

	req, err := http.NewRequest(http.MethodGet, uri, nil)
	suite.Nil(err)

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

	handler := http.HandlerFunc(handlers.Index)
	handler.ServeHTTP(suite.rr, req)
	exceptedBody, _ := json.Marshal(res.Data(userPaginate))
	suite.Equal(string(exceptedBody), suite.rr.Body.String())
	suite.Equal(http.StatusOK, suite.rr.Code)
}

func (suite *UserRouteTestSuite) TestIndexRouteOnFailed() {
	ucase := new(mocks.Usecase)
	handlers := route.NewHandler(ucase)

	defer ucase.AssertExpectations(suite.T())

	var page, limit uint
	page = 1
	limit = 10

	req, err := http.NewRequest(http.MethodGet, "/users", nil)
	suite.Nil(err)

	indexErr := errors.New("internal server error")
	ucase.On("Paginate", page, limit).Return(data.UserPaginate{}, indexErr).Once()

	handler := http.HandlerFunc(handlers.Index)
	handler.ServeHTTP(suite.rr, req)
	exceptedBody, _ := json.Marshal(res.Data(res.Error(indexErr.Error())))
	suite.Equal(string(exceptedBody), suite.rr.Body.String())
	suite.Equal(http.StatusInternalServerError, suite.rr.Code)
}
