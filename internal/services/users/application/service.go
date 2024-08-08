package application

import (
	"fmt"

	appError "banana-account-book.com/internal/libs/app-error"
	httpCode "banana-account-book.com/internal/libs/http/code"
	userModel "banana-account-book.com/internal/services/users/domain"
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

func (s *UserService) SignUp(email, password, name string) error {
	if _, ok, err := s.userRepository.FindByEmail(email); ok || err != nil {
		if ok {
			return appError.New(httpCode.Conflict, fmt.Sprintf("Email(%s) already exists", email), "Email already exists")
		}
		return appError.Wrap(err)
	}

	user, err := userModel.New(email, password, name, []string{"local"})
	if err != nil {
		return appError.Wrap(err)
	}
	return s.userRepository.Save(user)
}
