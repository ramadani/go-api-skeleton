package tests

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/ramadani/go-api-skeleton/app/user/data"
	"github.com/ramadani/go-api-skeleton/app/user/route"
	"github.com/ramadani/go-api-skeleton/app/user/usecase/mocks"
	hl "github.com/ramadani/go-api-skeleton/commons/handler"
)

func (suite *UserRouteTestSuite) TestStore() {
	ucase := new(mocks.Usecase)
	handlers := route.NewHandler(ucase)

	defer ucase.AssertExpectations(suite.T())

	input := url.Values{}
	input.Add("name", "FooBar")
	input.Add("email", "foo@example.com")
	input.Add("password", "secret")

	req, err := http.NewRequest(http.MethodPost, "/users", strings.NewReader(input.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(input.Encode())))
	suite.Nil(err)

	user := data.User{ID: 1, Name: "FooBar", Email: "foo@example.com"}
	ucase.On("Create", "FooBar", "foo@example.com", "secret").Return(user, nil).Once()

	handler := http.HandlerFunc(handlers.Store)
	handler.ServeHTTP(suite.rr, req)
	expectedBody, _ := json.Marshal(hl.ResponseData{Data: user})
	suite.Equal(string(expectedBody), suite.rr.Body.String())
	suite.Equal(http.StatusOK, suite.rr.Code)
}

func (suite *UserRouteTestSuite) TestStoreFailed() {
	ucase := new(mocks.Usecase)
	handlers := route.NewHandler(ucase)

	defer ucase.AssertExpectations(suite.T())

	input := url.Values{}
	input.Add("name", "FooBar")
	input.Add("email", "foo@example.com")
	input.Add("password", "secret")

	req, err := http.NewRequest(http.MethodPost, "/users", strings.NewReader(input.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(input.Encode())))
	suite.Nil(err)

	user := data.User{}
	createErr := errors.New("cannot create a new user")
	ucase.On("Create", "FooBar", "foo@example.com", "secret").Return(user, createErr).Once()

	handler := http.HandlerFunc(handlers.Store)
	handler.ServeHTTP(suite.rr, req)
	expectedBody, _ := json.Marshal(hl.ResponseData{Data: hl.ResponseError{Message: createErr.Error()}})
	suite.Equal(string(expectedBody), suite.rr.Body.String())
	suite.Equal(http.StatusInternalServerError, suite.rr.Code)
}
