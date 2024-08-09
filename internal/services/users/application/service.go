package application

import (
	"banana-account-book.com/internal/services/users/infrastructure"
)

type UserService struct {
	userRepository infrastructure.UserRepository
}

func NewUserService(userRepository infrastructure.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}
