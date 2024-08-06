package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"banana-account-book.com/internal/config"
	"banana-account-book.com/internal/libs/db"
	"banana-account-book.com/internal/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"go.uber.org/fx"
)

type App struct {
	*fiber.App
}

func New() *App {
	fiber := fiber.New()
	return &App{fiber}
}

func (app *App) Start(port string) error {
	return app.Listen(":" + port)
}

func (app *App) Stop() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		fmt.Println("Gracefully shutting down...")
		app.Shutdown()
	}()
}

func NewServer(lc fx.Lifecycle) *App {
	app := New()

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				db.Init()
				app.Get("/swagger/*", swagger.HandlerDefault)
				app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
					URL:         config.Origin + "/doc.json",
					DeepLinking: false,
					// Expand ("list") or Collapse ("none") tag groups by default
					DocExpansion: "none",
				}))
				router.Route(app.App)
				port := os.Getenv("PORT")
				fmt.Println("🔥Server started on port:", port, "🔥")
				app.Start(port)
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			app.Stop()
			return nil
		},
	})
	return app
}
