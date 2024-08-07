package domain

import (
	"banana-account-book.com/internal/libs/entity"
	"github.com/google/uuid"
)

type User struct {
	entity.SoftDeletableAggregate
	Id           uuid.UUID `json:"id" gorm:"primaryKey; type:uuid"`
	Email        string    `json:"email" gorm:"unique;type:varchar(50)"`
	Password     string    `json:"password" gorm:"type:varchar(255)"`
	Name         string    `json:"name" gorm:"type:varchar(50)"`
	Providers    []string  `json:"providers" gorm:"type:text[]"`
	RefreshToken string    `json:"refreshToken" gorm:"column:refreshToken;type:varchar(255)"`
}

func (u *User) TableName() string {
	return "user"
}
