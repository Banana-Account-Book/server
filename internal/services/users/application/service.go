package application

import (
	"banana-account-book.com/internal/services/users/domain"
	"banana-account-book.com/internal/services/users/dto"
	"banana-account-book.com/internal/services/users/infrastructure"
	"github.com/google/uuid"
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

func (s *UserService) Update(id uuid.UUID, body dto.UpdateUserRequestBody) (dto.UpdateUserResponse, error) {
	var (
		user *domain.User
		err  error
	)

	err = s.db.Transaction(func(tx *gorm.DB) error {
		user, err = s.userRepository.FindOneOrFail(tx, id)
		if err != nil {
			return err
		}

		user.Update(domain.UpdateType{Name: &body.Name})

		return s.userRepository.Save(tx, user)
	})

	return dto.UpdateUserResponse{Id: user.Id, Email: user.Email, Name: user.Name}, err
}
