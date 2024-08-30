package dto

import "banana-account-book.com/internal/types"

type CreateTransactionRequest struct {
	Title         string              `json:"title" validate:"required"`
	Description   string              `json:"description" `
	PeriodStartOn types.CalendarDate  `json:"periodStartOn" validate:"required,calendardate"`
	PeriodEndOn   *types.CalendarDate `json:"periodEndOn" validate:"omitempty,calendardate"`
	Type          string              `json:"type" validate:"required,oneof=income expense"`
	RepeatType    string              `json:"repeatType" validate:"required,oneof=none daily weekly monthly yearly"`
	Amount        int                 `json:"amount" validate:"required,gte=0"`
}
