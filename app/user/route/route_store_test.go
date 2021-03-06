package route_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/ramadani/go-api-skeleton/app/user/validators"

	"github.com/gorilla/mux"
	"github.com/ramadani/go-api-skeleton/app/user/data"
	"github.com/ramadani/go-api-skeleton/app/user/route"
	"github.com/ramadani/go-api-skeleton/app/user/usecase/mocks"
	"github.com/ramadani/go-api-skeleton/commons/http/res"
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

	router := mux.NewRouter()
	router.HandleFunc("/users", handlers.Store).Methods(http.MethodPost)
	router.ServeHTTP(suite.rr, req)

	expectedBody, _ := json.Marshal(res.Data(user))
	suite.Equal(string(expectedBody), suite.rr.Body.String())
	suite.Equal(http.StatusCreated, suite.rr.Code)
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

	router := mux.NewRouter()
	router.HandleFunc("/users", handlers.Store).Methods(http.MethodPost)
	router.ServeHTTP(suite.rr, req)

	expectedBody, _ := json.Marshal(res.Data(res.Message(createErr.Error())))
	suite.Equal(string(expectedBody), suite.rr.Body.String())
	suite.Equal(http.StatusInternalServerError, suite.rr.Code)
}

func (suite *UserRouteTestSuite) TestUserInputRequiredOnStore() {
	ucase := new(mocks.Usecase)
	handlers := route.NewHandler(ucase)

	defer ucase.AssertExpectations(suite.T())

	input := url.Values{}

	req, err := http.NewRequest(http.MethodPost, "/users", strings.NewReader(input.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(input.Encode())))
	suite.Nil(err)

	router := mux.NewRouter()
	router.HandleFunc("/users", handlers.Store).Methods(http.MethodPost)
	router.ServeHTTP(suite.rr, req)

	errs := validators.NewValidator().Validate(data.UserInput{})
	expectedBody, _ := json.Marshal(res.Data(res.ValidationError(errs)))
	suite.Equal(string(expectedBody), suite.rr.Body.String())
	suite.Equal(http.StatusBadRequest, suite.rr.Code)
}

func (suite *UserRouteTestSuite) TestEmailMustValidOnStore() {
	ucase := new(mocks.Usecase)
	handlers := route.NewHandler(ucase)

	defer ucase.AssertExpectations(suite.T())

	input := url.Values{}
	input.Add("name", "FooBar")
	input.Add("email", "foo")
	input.Add("password", "secret")

	req, err := http.NewRequest(http.MethodPost, "/users", strings.NewReader(input.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(input.Encode())))
	suite.Nil(err)

	router := mux.NewRouter()
	router.HandleFunc("/users", handlers.Store).Methods(http.MethodPost)
	router.ServeHTTP(suite.rr, req)

	errs := validators.NewValidator().Validate(data.UserInput{Name: "FooBar", Email: "foo", Password: "secret"})
	expectedBody, _ := json.Marshal(res.Data(res.ValidationError(errs)))
	suite.Equal(string(expectedBody), suite.rr.Body.String())
	suite.Equal(http.StatusBadRequest, suite.rr.Code)
}
