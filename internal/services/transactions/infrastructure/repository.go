package infrastructure

import (
	"fmt"

	appError "banana-account-book.com/internal/libs/app-error"
	httpCode "banana-account-book.com/internal/libs/http/code"
	"banana-account-book.com/internal/services/transactions/domain"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	Save(db *gorm.DB, Transaction *domain.Transaction) error
}

type TransactionRepositoryImpl struct {
	manager *gorm.DB
}

func NewTransactionRepository(manager *gorm.DB) TransactionRepository {
	return &TransactionRepositoryImpl{manager: manager}
}

func (r *TransactionRepositoryImpl) Save(db *gorm.DB, transaction *domain.Transaction) error {
	if db == nil {
		db = r.manager
	}

	if err := db.Session(&gorm.Session{FullSaveAssociations: true}).Save(transaction).Error; err != nil {
		return appError.New(httpCode.InternalServerError, fmt.Sprintf("Failed to save transaction. %s", err.Error()), "")
	}
	return nil
}
