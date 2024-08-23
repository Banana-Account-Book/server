package domain

import (
	"banana-account-book.com/internal/libs/ddd"
	"github.com/google/uuid"
)

type AccountBook struct {
	ddd.SoftDeletableAggregate
	Id     uuid.UUID `json:"id" gorm:"primaryKey; type:uuid; column:id"`
	UserId uuid.UUID `json:"userId" gorm:"type:uuid; column:userId; not null;"`
	Name   string    `json:"name" gorm:"type:varchar(50); column:name; not null;"`
}

func (a *AccountBook) TableName() string {
	return "account_book"
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
