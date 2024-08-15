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

	t.Run("HasProvider 테스트", func(t *testing.T) {
		t.Run("해당 provider를 가지고 있으면 true를 반환한다.", func(t *testing.T) {
			user, _ := domain.New("email", "name", []string{"kakao"})
			assert.True(t, user.HasProvider("kakao"))
		})

		t.Run("해당 provider를 가지고 있으면 false를 반환한다.", func(t *testing.T) {
			user, _ := domain.New("email", "name", []string{"kakao"})
			assert.False(t, user.HasProvider("naver"))
		})
	})

	t.Run("AddProvider 테스트", func(t *testing.T) {
		t.Run("provider를 추가한다.", func(t *testing.T) {
			user, _ := domain.New("email", "name", []string{"kakao"})
			user.AddProvider("naver")
			assert.Equal(t, user.Providers, pq.StringArray{"kakao", "naver"})
		})
	})
}
