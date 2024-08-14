package domain_test

import (
	"testing"

	"banana-account-book.com/internal/services/users/domain"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestAccount(t *testing.T) {
	t.Run("New 테스트", func(t *testing.T) {
		t.Run("email, name, provider를 받아 User 객체를 만든다", func(t *testing.T) {
			user, _ := domain.New("email", "name", []string{"kakao"})
			assert.Equal(t, user.Email, "email")
			assert.Equal(t, user.Name, "name")
			assert.Equal(t, user.Providers, pq.StringArray{"kakao"})
		})
	})
}
