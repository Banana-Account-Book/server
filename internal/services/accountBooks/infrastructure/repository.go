package infrastructure

import (
	"fmt"

	appError "banana-account-book.com/internal/libs/app-error"
	httpCode "banana-account-book.com/internal/libs/http/code"
	"banana-account-book.com/internal/services/accountBooks/domain"
	domainSpec "banana-account-book.com/internal/services/accountBooks/domain/specs"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AccountBookRepository interface {
	Save(db *gorm.DB, accountBook *domain.AccountBook) error
	FindByUserId(db *gorm.DB, userId uuid.UUID) ([]*domain.AccountBook, bool, error)
	FindSpec(db *gorm.DB, spec domainSpec.AccountBookSpec) ([]*domain.AccountBook, error)
	Delete(db *gorm.DB, accountBooks []*domain.AccountBook) error
}

type AccountBookRepositoryImpl struct {
	manager *gorm.DB
}

func NewAccountBookRepository(manager *gorm.DB) AccountBookRepository {
	return &AccountBookRepositoryImpl{manager: manager}
}

func (r *AccountBookRepositoryImpl) Save(db *gorm.DB, accountBook *domain.AccountBook) error {
	if db == nil {
		db = r.manager
	}

	if err := db.Save(accountBook).Error; err != nil {
		return appError.New(httpCode.InternalServerError, fmt.Sprintf("Failed to save account book. %s", err.Error()), "")
	}
	return nil
}

func (r *AccountBookRepositoryImpl) FindByUserId(db *gorm.DB, userId uuid.UUID) ([]*domain.AccountBook, bool, error) {
	if db == nil {
		db = r.manager
	}

	accountBooks := []*domain.AccountBook{}
	if err := db.Where("'userId' = ?::uuid", userId).Find(&accountBooks).Error; err != nil {
		return nil, false, appError.New(httpCode.InternalServerError, fmt.Sprintf("Failed to findByUserId account books. %s", err.Error()), "")
	}
	if len(accountBooks) == 0 {
		return nil, false, nil
	}
	return accountBooks, true, nil
}

func (r *AccountBookRepositoryImpl) FindSpec(db *gorm.DB, spec domainSpec.AccountBookSpec) ([]*domain.AccountBook, error) {
	if db == nil {
		db = r.manager
	}

	accountBooks, err := spec.Find(db)
	if err != nil {
		return nil, appError.Wrap(err)
	}

	return accountBooks, nil
}

func (r *AccountBookRepositoryImpl) Delete(db *gorm.DB, accountBooks []*domain.AccountBook) error {
	if db == nil {
		db = r.manager
	}

	if err := db.Delete(accountBooks).Error; err != nil {
		return appError.New(httpCode.InternalServerError, fmt.Sprintf("Failed to delete account books. %s", err.Error()), "")
	}
	return nil
}
