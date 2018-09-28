package route_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ramadani/go-api-skeleton/app/user/data"
	"github.com/ramadani/go-api-skeleton/app/user/route"
	"github.com/ramadani/go-api-skeleton/app/user/usecase/mocks"
	"github.com/ramadani/go-api-skeleton/commons/http/res"
)

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

	expectedBody, _ := json.Marshal(res.Data(res.Message(findErr.Error())))
	suite.Equal(string(expectedBody), suite.rr.Body.String())
	suite.Equal(http.StatusNotFound, suite.rr.Code)
}
