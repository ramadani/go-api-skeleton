package main

import (
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	userRoute "github.com/ramadani/go-api-skeleton/app/user/route"
	welcomeRoute "github.com/ramadani/go-api-skeleton/app/welcome/route"
	"github.com/ramadani/go-api-skeleton/bootstrap"
)

type example struct {
	Name string `json:"name"`
}

func main() {
	r := mux.NewRouter()

	// Our routes
	welcomeRoute.New(r)
	userRoute.New(r)

	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	compressedRouter := handlers.CompressHandler(loggedRouter)

	app := bootstrap.New(compressedRouter)
	app.Run(3000)
}
