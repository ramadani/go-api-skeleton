package bootstrap

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/ramadani/go-api-skeleton/routes"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

type App struct {
	e *echo.Echo
}

func (app App) InitMiddleware() {
	app.e.Use(middleware.Logger())
	app.e.Use(middleware.Recover())
}

func (app App) InitRoutes() {
	routes.NewApiRoutes(app.e)
}

func (app App) Run(port int) {
	app.e.Logger.SetLevel(log.INFO)

	go func() {
		if err := app.e.Start(fmt.Sprintf(":%d", port)); err != nil {
			app.e.Logger.Info("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := app.e.Shutdown(ctx); err != nil {
		app.e.Logger.Fatal(err)
	}
}

func New(e *echo.Echo) *App {
	return &App{e}
}
