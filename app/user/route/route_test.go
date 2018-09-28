package route_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"

	"github.com/gorilla/mux"

	"github.com/ramadani/go-api-skeleton/app/user/data"
	"github.com/ramadani/go-api-skeleton/app/user/route"
	"github.com/ramadani/go-api-skeleton/app/user/usecase/mocks"
	"github.com/ramadani/go-api-skeleton/commons/http/res"
	"github.com/stretchr/testify/suite"
)

type UserRouteTestSuite struct {
	suite.Suite
	rr *httptest.ResponseRecorder
}

func (suite *UserRouteTestSuite) SetupTest() {
	suite.rr = httptest.NewRecorder()
}

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
	expectedBody, _ := json.Marshal(res.Data(user))
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
	expectedBody, _ := json.Marshal(res.Data(res.Error(createErr.Error())))
	suite.Equal(string(expectedBody), suite.rr.Body.String())
	suite.Equal(http.StatusInternalServerError, suite.rr.Code)
}

func (suite *UserRouteTestSuite) TestFind() {
	ucase := new(mocks.Usecase)
	handlers := route.NewHandler(ucase)

	defer ucase.AssertExpectations(suite.T())

	id := uint(1)
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/users/%d", id), nil)
	suite.Nil(err)

	user := data.User{ID: id, Name: "FooBar", Email: "foo@example.com"}
	ucase.On("FindByID", id).Return(user, nil).Once()

	router := mux.NewRouter()
	router.HandleFunc("/users/{id:[0-9]+}", handlers.Find).Methods(http.MethodGet)
	router.ServeHTTP(suite.rr, req)

	expectedBody, _ := json.Marshal(res.Data(user))
	suite.Equal(string(expectedBody), suite.rr.Body.String())
	suite.Equal(http.StatusOK, suite.rr.Code)
}

func (suite *UserRouteTestSuite) TestFindNotFound() {
	ucase := new(mocks.Usecase)
	handlers := route.NewHandler(ucase)

	defer ucase.AssertExpectations(suite.T())

	id := uint(1)
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/users/%d", id), nil)
	suite.Nil(err)

	user := data.User{}
	findErr := errors.New("not found")
	ucase.On("FindByID", id).Return(user, findErr).Once()

	router := mux.NewRouter()
	router.HandleFunc("/users/{id:[0-9]+}", handlers.Find).Methods(http.MethodGet)
	router.ServeHTTP(suite.rr, req)

	expectedBody, _ := json.Marshal(res.Data(res.Error(findErr.Error())))
	suite.Equal(string(expectedBody), suite.rr.Body.String())
	suite.Equal(http.StatusNotFound, suite.rr.Code)
}

func TestUserRouteTestSuite(t *testing.T) {
	suite.Run(t, new(UserRouteTestSuite))
}
