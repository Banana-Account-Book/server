package application

import (
	"banana-account-book.com/internal/services/users/infrastructure"
	"gorm.io/gorm"
)

type UserService struct {
	userRepository infrastructure.UserRepository
	db             *gorm.DB
}

func NewUserService(userRepository infrastructure.UserRepository, db *gorm.DB) *UserService {
	return &UserService{
		userRepository: userRepository,
		db:             db,
	}
}
