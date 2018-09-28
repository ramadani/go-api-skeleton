package route_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ramadani/go-api-skeleton/app/user/route"
	"github.com/ramadani/go-api-skeleton/app/user/usecase/mocks"
	"github.com/ramadani/go-api-skeleton/commons/http/res"
)

func (suite *UserRouteTestSuite) TestDeleteFailed() {
	ucase := new(mocks.Usecase)
	handlers := route.NewHandler(ucase)

	defer ucase.AssertExpectations(suite.T())

	id := uint(1)
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("/users/%d", id), nil)
	suite.Nil(err)

	findErr := errors.New("not found")
	ucase.On("Delete", id).Return(findErr).Once()

	router := mux.NewRouter()
	router.HandleFunc("/users/{id:[0-9]+}", handlers.Delete).Methods(http.MethodDelete)
	router.ServeHTTP(suite.rr, req)

	expectedBody, _ := json.Marshal(res.Data(res.Message(findErr.Error())))
	suite.Equal(string(expectedBody), suite.rr.Body.String())
	suite.Equal(http.StatusNotFound, suite.rr.Code)
}

func (suite *UserRouteTestSuite) TestDeleteSuccess() {
	ucase := new(mocks.Usecase)
	handlers := route.NewHandler(ucase)

	defer ucase.AssertExpectations(suite.T())

	id := uint(1)
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("/users/%d", id), nil)
	suite.Nil(err)

	ucase.On("Delete", id).Return(nil).Once()

	router := mux.NewRouter()
	router.HandleFunc("/users/{id:[0-9]+}", handlers.Delete).Methods(http.MethodDelete)
	router.ServeHTTP(suite.rr, req)

	expectedBody, _ := json.Marshal(res.Data(res.Message("user has been deleted")))
	suite.Equal(string(expectedBody), suite.rr.Body.String())
	suite.Equal(http.StatusOK, suite.rr.Code)
}
