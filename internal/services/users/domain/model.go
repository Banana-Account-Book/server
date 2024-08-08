package domain

import (
	"strconv"

	"banana-account-book.com/internal/config"
	appError "banana-account-book.com/internal/libs/app-error"
	"banana-account-book.com/internal/libs/entity"
	httpCode "banana-account-book.com/internal/libs/http/code"
	"github.com/google/uuid"
	"github.com/lib/pq"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	entity.SoftDeletableAggregate
	Id           uuid.UUID      `json:"id" gorm:"primaryKey; type:uuid"`
	Email        string         `json:"email" gorm:"unique;type:varchar(50)"`
	Password     string         `json:"password" gorm:"type:varchar(255)"`
	Name         string         `json:"name" gorm:"type:varchar(50)"`
	Providers    pq.StringArray `json:"providers" gorm:"type:text[];"`
	RefreshToken string         `json:"refreshToken" gorm:"column:refreshToken;type:varchar(255)"`
}

func (u *User) TableName() string {
	return "user"
}

func New(email, password, name string, providers []string) (*User, error) {
	uuId, err := uuid.NewV7()
	if err != nil {
		return nil, appError.New(httpCode.InternalServerError, "Failed to create new user. Can not generate uuid.", "Internal Server Error")
	}

	salt, err := strconv.Atoi(config.Salt)
	if err != nil {
		return nil, appError.New(httpCode.InternalServerError, "Failed to create new user. Can not convert salt to int.", "Internal Server Error")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), salt)
	if err != nil {
		return nil, appError.New(httpCode.InternalServerError, "Failed to create new user. Can not hash password.", "Internal Server Error")
	}

	user := &User{
		Id:        uuId,
		Email:     email,
		Password:  string(hashedPassword),
		Name:      name,
		Providers: pq.StringArray(providers),
	}

	return user, nil
}
