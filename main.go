package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	userRoute "github.com/ramadani/go-api-skeleton/app/user/route"
	welcomeRoute "github.com/ramadani/go-api-skeleton/app/welcome/route"
	"github.com/ramadani/go-api-skeleton/bootstrap"
)

func main() {
	var handler http.Handler

	r := mux.NewRouter()
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/go_api")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// Our routes
	welcomeRoute.New(r)
	userRoute.New(r, db)

	handler = r
	handler = handlers.LoggingHandler(os.Stdout, handler)
	handler = handlers.CompressHandler(handler)

	app := bootstrap.New(handler)
	app.Run(3000)
}
