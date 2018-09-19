package route

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ramadani/go-api-skeleton/app/user/repository"
)

// New user routes
func New(router *mux.Router, db *sql.DB) {
	repository.NewMySQLRepository(db)
	handler := NewHandler()

	router.HandleFunc("/users", handler.Index()).Methods(http.MethodGet)
}
