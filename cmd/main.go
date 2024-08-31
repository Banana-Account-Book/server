package main

import (
	"banana-account-book.com/internal/app"
	"banana-account-book.com/internal/libs/db"
	"banana-account-book.com/internal/libs/validate"
	"banana-account-book.com/internal/middlewares"
	"banana-account-book.com/internal/services/accountBooks"
	"banana-account-book.com/internal/services/auth"
	"banana-account-book.com/internal/services/roles"
	"banana-account-book.com/internal/services/transactions"
	"banana-account-book.com/internal/services/users"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

func main() {
	validate.Init()
	fx.New(
		fx.Provide(db.Init),
		fx.Provide(app.NewServer),
		fx.Provide(middlewares.NewAuthHandler),
		users.Module,
		auth.Module,
		roles.Module,
		accountBooks.Module,
		transactions.Module,
		fx.Invoke(func(*app.App) {}),
		fx.Invoke(func(*gorm.DB) {}),
	).Run()
}
