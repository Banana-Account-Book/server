package domain

import (
	"banana-account-book.com/internal/libs/ddd"
	"github.com/google/uuid"
)

type AccountBook struct {
	ddd.SoftDeletableAggregate
	Id     uuid.UUID `json:"id" gorm:"primaryKey; type:uuid; column:id"`
	UserId uuid.UUID `json:"userId" gorm:"type:uuid; column:userId"`
	Name   string    `json:"name" gorm:"type:varchar(50); column:name"`
}

func (a *AccountBook) TableName() string {
	return "account"
}

func New(userId uuid.UUID, name string) (*AccountBook, error) {
	uuId, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	accountBook := &AccountBook{
		Id:     uuId,
		UserId: userId,
		Name:   name,
	}

	return accountBook, nil
}
