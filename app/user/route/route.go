package route

import (
	"net/http"

	"github.com/gorilla/mux"
)

// New user routes
func New(router *mux.Router) {
	handler := NewHandler()

	router.HandleFunc("/users", handler.Index).Methods(http.MethodGet)
}
