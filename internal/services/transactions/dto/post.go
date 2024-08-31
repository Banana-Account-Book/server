package dto

import (
	"banana-account-book.com/internal/services/transactions/domain"
	"banana-account-book.com/internal/types"
)

type CreateTransactionRequest struct {
	Title         string                 `json:"title" validate:"required"`
	Description   string                 `json:"description" `
	PeriodStartOn types.CalendarDate     `json:"periodStartOn" validate:"required,calendardate"`
	PeriodEndOn   *types.CalendarDate    `json:"periodEndOn" validate:"omitempty,calendardate"`
	Type          domain.TransactionType `json:"type" validate:"required,oneof=Income Expense Transfer"`
	RepeatType    domain.RepeatPeriod    `json:"repeatType" validate:"required,oneof=None Daily Weekly Monthly Yearly Custom"`
	Amount        int                    `json:"amount" validate:"required,gte=0"`
}
