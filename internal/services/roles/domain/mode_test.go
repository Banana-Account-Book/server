package domain_test

import (
	"testing"

	"banana-account-book.com/internal/services/roles/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestRole(t *testing.T) {
	t.Run("New 테스트", func(t *testing.T) {
		t.Run("accountBookId, userId, type을 받아 Role 객체를 만든다", func(t *testing.T) {
			id := "11111111-1111-1111-1111-111111111111"
			testId := uuid.MustParse(id)
			role, _ := domain.New(testId, testId, "owner")
			assert.Equal(t, role.UserId, uuid.MustParse(id))
			assert.Equal(t, role.AccountBookId, uuid.MustParse(id))
			assert.Equal(t, role.Type, "owner")
		})
		t.Run("type은 owner나 member여야 한다", func(t *testing.T) {
			t.Run("owner", func(t *testing.T) {
				role, _ := domain.New(uuid.New(), uuid.New(), "owner")
				assert.Equal(t, role.Type, "owner")
			})
			t.Run("member", func(t *testing.T) {
				role, _ := domain.New(uuid.New(), uuid.New(), "member")
				assert.Equal(t, role.Type, "member")
			})
		})
		t.Run("type이 owner나 member가 아니면 에러를 반환한다", func(t *testing.T) {
			_, err := domain.New(uuid.New(), uuid.New(), "test")
			assert.Error(t, err)
			assert.Equal(t, err.Error(), "Invalid role type: test")
		})
	})
}
