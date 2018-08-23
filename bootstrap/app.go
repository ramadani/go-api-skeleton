package bootstrap

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/spf13/viper"

	"github.com/ramadani/go-api-skeleton/provider"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

type App struct {
	fw  *echo.Echo
	cog *viper.Viper
}

func (app App) Run() {
	bootables := []Bootable{
		provider.InitMiddleware(app.fw),
		provider.InitRoute(app.fw),
	}

	for _, bootable := range bootables {
		bootable.Boot()
	}

	port := app.cog.GetInt("port")
	app.fw.Logger.SetLevel(log.INFO)

	go func() {
		if err := app.fw.Start(fmt.Sprintf(":%d", port)); err != nil {
			app.fw.Logger.Info("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := app.fw.Shutdown(ctx); err != nil {
		app.fw.Logger.Fatal(err)
	}
}

func New(fw *echo.Echo, cog *viper.Viper) *App {
	return &App{fw, cog}
}
