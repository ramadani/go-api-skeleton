package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	userRoute "github.com/ramadani/go-api-skeleton/app/user/route"
	welcomeRoute "github.com/ramadani/go-api-skeleton/app/welcome/route"
	"github.com/ramadani/go-api-skeleton/bootstrap"
)

func main() {
	r := mux.NewRouter()
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/go_api")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// Our routes
	welcomeRoute.New(r)
	userRoute.New(r, db)

	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	compressedRouter := handlers.CompressHandler(loggedRouter)

	app := bootstrap.New(compressedRouter)
	app.Run(3000)
}
