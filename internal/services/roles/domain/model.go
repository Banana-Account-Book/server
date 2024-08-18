package domain

import (
	"fmt"

	appError "banana-account-book.com/internal/libs/app-error"
	"banana-account-book.com/internal/libs/ddd"
	httpCode "banana-account-book.com/internal/libs/http/code"
	"github.com/google/uuid"
)

type Role struct {
	ddd.SoftDeletableAggregate
	Id            int       `json:"id" gorm:"primaryKey; autoIncrement"`
	Type          string    `json:"type" gorm:"type:varchar(20); not null;"`
	AccountBookId uuid.UUID `json:"accountBookId" gorm:"type:uuid; column:accountBookId; not null;"`
	UserId        uuid.UUID `json:"userId" gorm:"type:uuid;column:userId; not null;"`
}

func (r *Role) TableName() string {
	return "role"
}

func New(userId, accountBookId uuid.UUID, roleType string) (*Role, error) {
	if roleType != "owner" && roleType != "member" {
		message := fmt.Sprintf("Invalid role type: %s", roleType)
		return nil, appError.New(httpCode.BadRequest, message, message)
	}

	role := &Role{
		Type:          roleType,
		AccountBookId: accountBookId,
		UserId:        userId,
	}

	return role, nil
}
