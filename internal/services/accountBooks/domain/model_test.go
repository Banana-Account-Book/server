package domain_test

import (
	"testing"

	"banana-account-book.com/internal/services/accountBooks/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAccountBook(t *testing.T) {
	t.Run("New 테스트", func(t *testing.T) {
		t.Run("userId, name을 받아 AccountBook 객체를 만든다", func(t *testing.T) {
			testId := uuid.MustParse("11111111-1111-1111-1111-111111111111")
			accountBook, _ := domain.New(testId, "testName")
			assert.Equal(t, accountBook.UserId, uuid.MustParse("11111111-1111-1111-1111-111111111111"))
			assert.Equal(t, accountBook.Name, "testName")
		})
	})
}
