package application

import (
	"fmt"

	appError "banana-account-book.com/internal/libs/app-error"
	"banana-account-book.com/internal/services/accountBooks/domain"
	domainSpec "banana-account-book.com/internal/services/accountBooks/domain/specs"
	"banana-account-book.com/internal/services/accountBooks/infrastructure"
	roleModel "banana-account-book.com/internal/services/roles/domain"
	roleInfra "banana-account-book.com/internal/services/roles/infrastructure"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AccountBookService struct {
	accountBookRepository infrastructure.AccountBookRepository
	roleRepository        roleInfra.RoleRepository
	db                    *gorm.DB
}

func NewAccountBookService(accountBookRepository infrastructure.AccountBookRepository, db *gorm.DB, roleRepository roleInfra.RoleRepository) *AccountBookService {
	return &AccountBookService{
		accountBookRepository: accountBookRepository,
		roleRepository:        roleRepository,
		db:                    db,
	}
}

func (s *AccountBookService) Add(userId uuid.UUID, name string) error {
	err := s.db.Transaction(func(tx *gorm.DB) error {
		accountBook, err := domain.New(userId, name)
		if err != nil {
			return appError.Wrap(err)
		}

		err = s.accountBookRepository.Save(tx, accountBook)
		if err != nil {
			return appError.Wrap(err)
		}

		role, err := roleModel.New(userId, accountBook.Id, "owner")
		if err != nil {
			return appError.Wrap(err)
		}
		fmt.Println("!!!", s.roleRepository)

		if err := s.roleRepository.Save(tx, role); err != nil {
			return appError.Wrap(err)
		}

		return nil
	})

	return err
}

func (s *AccountBookService) Delete(roles []*roleModel.Role, accountId uuid.UUID) error {
	err := s.db.Transaction(func(tx *gorm.DB) error {
		accountBooks, err := s.accountBookRepository.FindSpec(tx, domainSpec.NewDeletableAccountBookSpec(roles, accountId))

		if err != nil {
			return err
		}

		return s.accountBookRepository.Delete(tx, accountBooks)
	})

	if err != nil {
		return appError.Wrap(err)
	}

	return nil
}
