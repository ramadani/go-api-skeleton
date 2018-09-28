package route_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/ramadani/go-api-skeleton/app/user/data"
	"github.com/ramadani/go-api-skeleton/app/user/route"
	"github.com/ramadani/go-api-skeleton/app/user/usecase/mocks"
	"github.com/ramadani/go-api-skeleton/app/user/validators"
	"github.com/ramadani/go-api-skeleton/commons/http/res"
)

func (suite *UserRouteTestSuite) TestNameRequiredOnUpdate() {
	ucase := new(mocks.Usecase)
	handlers := route.NewHandler(ucase)

	defer ucase.AssertExpectations(suite.T())

	id := uint(1)
	input := url.Values{}

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("/users/%d", id), strings.NewReader(input.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(input.Encode())))
	suite.Nil(err)

	router := mux.NewRouter()
	router.HandleFunc("/users/{id:[0-9]+}", handlers.Update).Methods(http.MethodPut)
	router.ServeHTTP(suite.rr, req)

	errs := validators.NewValidator().Validate(data.UserUpdateInput{})
	expectedBody, _ := json.Marshal(res.Data(res.ValidationError(errs)))
	suite.Equal(string(expectedBody), suite.rr.Body.String())
	suite.Equal(http.StatusBadRequest, suite.rr.Code)
}

func (suite *UserRouteTestSuite) TestUpdateFailed() {
	ucase := new(mocks.Usecase)
	handlers := route.NewHandler(ucase)

	defer ucase.AssertExpectations(suite.T())

	id := uint(1)
	input := url.Values{}
	input.Add("name", "BarFoo")

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("/users/%d", id), strings.NewReader(input.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(input.Encode())))
	suite.Nil(err)

	user := data.User{}
	updateErr := errors.New("cannot update user")
	ucase.On("Update", "BarFoo", id).Return(user, updateErr).Once()

	router := mux.NewRouter()
	router.HandleFunc("/users/{id:[0-9]+}", handlers.Update).Methods(http.MethodPut)
	router.ServeHTTP(suite.rr, req)

	expectedBody, _ := json.Marshal(res.Data(res.Error(updateErr.Error())))
	suite.Equal(string(expectedBody), suite.rr.Body.String())
	suite.Equal(http.StatusInternalServerError, suite.rr.Code)
}

func (suite *UserRouteTestSuite) TestUpdateSuccess() {
	ucase := new(mocks.Usecase)
	handlers := route.NewHandler(ucase)

	defer ucase.AssertExpectations(suite.T())

	id := uint(1)
	input := url.Values{}
	input.Add("name", "BarFoo")

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("/users/%d", id), strings.NewReader(input.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(input.Encode())))
	suite.Nil(err)

	user := data.User{ID: id, Name: "BarFoo", Email: "foo@example.com"}
	ucase.On("Update", "BarFoo", id).Return(user, nil).Once()

	router := mux.NewRouter()
	router.HandleFunc("/users/{id:[0-9]+}", handlers.Update).Methods(http.MethodPut)
	router.ServeHTTP(suite.rr, req)

	expectedBody, _ := json.Marshal(res.Data(user))
	suite.Equal(string(expectedBody), suite.rr.Body.String())
	suite.Equal(http.StatusOK, suite.rr.Code)
}
