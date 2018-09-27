package route

import (
	"database/sql"
	"net/http"

	"github.com/ramadani/go-api-skeleton/app/user/usecase"

	"github.com/gorilla/mux"
	"github.com/ramadani/go-api-skeleton/app/user/repository"
)

// New user routes
func New(router *mux.Router, db *sql.DB) {
	repo := repository.NewMySQLRepository(db)
	ucase := usecase.New(repo)
	handler := NewHandler(ucase)

	router.HandleFunc("/users", handler.Index()).Methods(http.MethodGet)
	router.HandleFunc("/users", handler.Store()).Methods(http.MethodPost)
}
