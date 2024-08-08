package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"banana-account-book.com/internal/config"
	"banana-account-book.com/internal/middlewares"
	"banana-account-book.com/internal/router"
	user "banana-account-book.com/internal/services/users/presentation"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/swagger"
	"go.uber.org/fx"
)

type App struct {
	*fiber.App
}

func New() *App {
	fiber := fiber.New(fiber.Config{
		ErrorHandler: middlewares.ErrorHandler,
	})
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

func NewServer(lc fx.Lifecycle, userController *user.UserController) *App {
	app := New()

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				// request logger
				app.Use(requestid.New(), logger.New(logger.Config{
					Format:     "${time} | ${pid} | ${locals:requestid} | ${status} - ${method} ${path}\u200b\n",
					TimeFormat: "2006-01-02 15:04:05",
					TimeZone:   "UTC",
				}))

				// swagger
				app.Get("/swagger/*", swagger.HandlerDefault, swagger.New(swagger.Config{ // custom
					URL:         config.Origin + "/doc.json",
					DeepLinking: false,
					// Expand ("list") or Collapse ("none") tag groups by default
					DocExpansion: "none",
				}))

				router.Route(app.App, userController)
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
