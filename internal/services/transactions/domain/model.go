package domain

import (
	"fmt"
	"time"

	appError "banana-account-book.com/internal/libs/app-error"
	"banana-account-book.com/internal/libs/ddd"
	httpCode "banana-account-book.com/internal/libs/http/code"
	"banana-account-book.com/internal/types"
	"github.com/google/uuid"
)

type TransactionType string

const (
	Income   TransactionType = "Income"
	Expense  TransactionType = "Expense"
	Transfer TransactionType = "Transfer"
)

type RepeatPeriod string

const (
	None    RepeatPeriod = "None"
	Daily   RepeatPeriod = "Daily"
	Weekly  RepeatPeriod = "Weekly"
	Monthly RepeatPeriod = "Monthly"
	Yearly  RepeatPeriod = "Yearly"
	Custom  RepeatPeriod = "Custom"
)

type TransactionDetails struct {
	UserId        uuid.UUID
	AccountBookId uuid.UUID
	Title         string
	Description   string
	PeriodStartOn types.CalendarDate
	PeriodEndOn   *types.CalendarDate
	Type          TransactionType
	RepeatType    RepeatPeriod
	Amount        int
}

type Transaction struct {
	ddd.SoftDeletableAggregate
	Id            uuid.UUID           `json:"id" gorm:"primaryKey; type:uuid; column:id"`
	AccountBookId uuid.UUID           `json:"accountBookId" gorm:"type:uuid; column:accountBookId; not null;"`
	UserId        uuid.UUID           `json:"userId" gorm:"type:uuid; column:userId; not null;"`
	Title         string              `json:"title" gorm:"type:varchar(50); column:title; not null;"`
	Description   string              `json:"description" gorm:"type:varchar(255); column:description;"`
	RegisteredAt  time.Time           `json:"registeredAt" gorm:"column:registeredAt; not null; type:timestamptz;"`
	PeriodStartOn types.CalendarDate  `json:"periodStart" gorm:"column:periodStartOn; not null; type:date;"`
	PeriodEndOn   *types.CalendarDate `json:"periodEnd" gorm:"column:periodEndOn; type:date;"`
	Type          TransactionType     `json:"type" gorm:"type:varchar(20); column:type; not null;"`
	Amount        int                 `json:"amount" gorm:"column:amount; not null;"`
	RepeatType    RepeatPeriod        `json:"repeatType" gorm:"type:varchar(20); column:repeatType;"`
	Exclusives    []Exclusive         `gorm:"foreignKey:TransactionId; references:Id"`
}

type Exclusive struct {
	Id            int                 `json:"id" gorm:"primaryKey; column:id; autoIncrement;"`
	UserId        uuid.UUID           `json:"userId" gorm:"type:uuid; column:userId; not null;"`
	PeriodStartOn types.CalendarDate  `json:"periodStart" gorm:"column:periodStartOn; not null; type:date;"`
	PeriodEndOn   *types.CalendarDate `json:"periodEnd" gorm:"column:periodEndOn; type:date;"`
	Title         string              `json:"title" gorm:"type:varchar(50); column:title; not null;"`
	Description   string              `json:"description" gorm:"type:varchar(255); column:description;"`
	Amount        int                 `json:"amount" gorm:"column:amount; not null;"`
	TransactionId uuid.UUID           `gorm:"type:uuid; column:transactionId; not null;"` // 외래키
	Transaction   Transaction         `gorm:"foreignKey:TransactionId; references:Id"`
}

func (t *Transaction) TableName() string {
	return "transaction"
}

func (e *Exclusive) TableName() string {
	return "exclusive"
}

func New(args TransactionDetails) (*Transaction, error) {
	uuId, err := uuid.NewV7()
	if err != nil {
		return nil, appError.New(httpCode.InternalServerError, fmt.Sprintf("failed to generate uuid: %v", err), "")
	}

	transaction := &Transaction{
		Id:            uuId,
		UserId:        args.UserId,
		AccountBookId: args.AccountBookId,
		Title:         args.Title,
		Description:   args.Description,
		RegisteredAt:  time.Now(),
		PeriodStartOn: args.PeriodStartOn,
		PeriodEndOn:   args.PeriodEndOn,
		Type:          args.Type,
		Amount:        args.Amount,
		RepeatType:    args.RepeatType,
		Exclusives:    []Exclusive{},
	}

	return transaction, nil
}
