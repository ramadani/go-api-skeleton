package route

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/ramadani/go-api-skeleton/app/user/usecase"

	"github.com/gorilla/mux"
	"github.com/ramadani/go-api-skeleton/app/user/repository"
)

// New user routes
func New(router *mux.Router, db *sql.DB) {
	repo := repository.NewMySQLRepository(db)
	ucase := usecase.New(repo)
	handler := NewHandler(ucase)

	router.HandleFunc("/users", handler.Index).Methods(http.MethodGet)
	router.HandleFunc("/users", handler.Store).Methods(http.MethodPost)
	router.HandleFunc("/users/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(mux.Vars(r)["id"])
		handler.Find(w, r, uint(id))
	}).Methods(http.MethodGet)
}
