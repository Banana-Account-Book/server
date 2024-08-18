package application

import (
	appError "banana-account-book.com/internal/libs/app-error"
	"banana-account-book.com/internal/services/accountBooks/domain"
	"banana-account-book.com/internal/services/accountBooks/infrastructure"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AccountBookService struct {
	accountBookRepository infrastructure.AccountBookRepository
	db                    *gorm.DB
}

func NewAccountBookService(accountBookRepository infrastructure.AccountBookRepository, db *gorm.DB) *AccountBookService {
	return &AccountBookService{
		accountBookRepository: accountBookRepository,
		db:                    db,
	}
}

func (s *AccountBookService) Add(ownerId uuid.UUID, name string) error {
	accountBook, err := domain.New(ownerId, name)
	if err != nil {
		return appError.Wrap(err)
	}

	err = s.accountBookRepository.Save(nil, accountBook)
	if err != nil {
		return appError.Wrap(err)
	}

	return nil
}
