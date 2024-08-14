package infrastructure

import (
	"fmt"

	appError "banana-account-book.com/internal/libs/app-error"
	httpCode "banana-account-book.com/internal/libs/http/code"
	"banana-account-book.com/internal/services/users/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindByEmail(db *gorm.DB, email string) (*domain.User, bool, error)
	Save(db *gorm.DB, user *domain.User) error
}

type UserRepositoryImpl struct {
	manager *gorm.DB
}

func NewRepository(manager *gorm.DB) UserRepository {
	return &UserRepositoryImpl{manager: manager}
}

func (r *UserRepositoryImpl) FindByEmail(db *gorm.DB, email string) (*domain.User, bool, error) {
	if db == nil {
		db = r.manager
	}

	users := []domain.User{}
	if err := db.Where("email = ?", email).Find(&users).Error; err != nil {
		return nil, false, appError.New(httpCode.InternalServerError, fmt.Sprintf("Failed to findByEmail user. %s", err.Error()), "")
	}
	if len(users) == 0 {
		return nil, false, nil
	}
	return &users[0], true, nil
}

func (r *UserRepositoryImpl) Save(db *gorm.DB, user *domain.User) error {
	if db == nil {
		db = r.manager
	}

	if err := db.Save(user).Error; err != nil {
		return appError.New(httpCode.InternalServerError, fmt.Sprintf("Failed to save user. %s", err.Error()), "")
	}
	return nil
}
