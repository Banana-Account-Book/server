package domain_test

import (
	"testing"

	"banana-account-book.com/internal/services/transactions/domain"
	"banana-account-book.com/internal/types"
	"github.com/google/uuid"
	"gopkg.in/go-playground/assert.v1"
)

func TestUser(t *testing.T) {
	t.Run("New 테스트", func(t *testing.T) {
		t.Run("정보를 받아 Transaction 객체를 만든다", func(t *testing.T) {
			transaction, _ := domain.New(domain.TransactionDetails{
				UserId:        uuid.MustParse("11111111-1111-1111-1111-111111111111"),
				AccountBookId: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
				Title:         "title",
				Description:   "description",
				PeriodStartOn: types.CalendarDate("2024-09-01"),
				PeriodEndOn:   nil,
				Type:          "income",
				RepeatType:    "none",
				Amount:        10000,
			})
			assert.Equal(t, transaction.Title, "title")
			assert.Equal(t, transaction.Description, "description")
			assert.Equal(t, transaction.PeriodStartOn, types.CalendarDate("2024-09-01"))
		})
	})
}
