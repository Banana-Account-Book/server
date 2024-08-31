package transactions

import (
	"banana-account-book.com/internal/services/transactions/application"
	"banana-account-book.com/internal/services/transactions/infrastructure"
	"banana-account-book.com/internal/services/transactions/presentation"
	"go.uber.org/fx"
)

var Module = fx.Module("transactions",
	fx.Provide(infrastructure.NewTransactionRepository),
	fx.Provide(application.NewTransactionService),
	fx.Provide(presentation.NewTransactionController),
)
