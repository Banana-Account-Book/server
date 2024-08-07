package domain

import "banana-account-book.com/internal/libs/entity"

type User struct {
	entity.SoftDeletableAggregate
	Id           string   `json:"id" gorm:"primaryKey; type:varchar(16)"`
	Email        string   `json:"email" gorm:"unique;type:varchar(50)"`
	Password     string   `json:"password" gorm:"type:varchar(255)"`
	Name         string   `json:"name" gorm:"type:varchar(50)"`
	Providers    []string `json:"providers" gorm:"type:jsonb"`
	RefreshToken string   `json:"refreshToken" gorm:"column:refreshToken;type:varchar(255)"`
}

func (u *User) TableName() string {
	return "user"
}
