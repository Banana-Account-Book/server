package accountBooks

import (
	"banana-account-book.com/internal/services/accountBooks/application"
	"banana-account-book.com/internal/services/accountBooks/infrastructure"
	"banana-account-book.com/internal/services/accountBooks/presentation"
	"go.uber.org/fx"
)

var Module = fx.Module("account-books",
	fx.Provide(application.NewAccountBookService),
	fx.Provide(infrastructure.NewAccountBookRepository),
	fx.Provide(presentation.NewAccountBookController),
)
