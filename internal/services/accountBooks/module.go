package accountBooks

import (
	"banana-account-book.com/internal/services/accountBooks/application"
	"banana-account-book.com/internal/services/accountBooks/infrastructure"
	"go.uber.org/fx"
)

var Module = fx.Module("account",
	fx.Provide(application.NewAccountBookService),
	fx.Provide(infrastructure.NewAccountBookRepository),
)
