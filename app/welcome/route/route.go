package route

import (
	"net/http"

	"github.com/gorilla/mux"
)

// New welcome routes
func New(router *mux.Router) {
	handler := NewHandler()

	router.HandleFunc("/", handler.Index).Methods(http.MethodGet)
}
