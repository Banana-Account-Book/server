package domain_specs_test

import (
	"fmt"
	"testing"
	"time"

	accountBookModel "banana-account-book.com/internal/services/accountBooks/domain"
	domain_specs "banana-account-book.com/internal/services/accountBooks/domain/specs"
	roleModel "banana-account-book.com/internal/services/roles/domain"
	"banana-account-book.com/internal/test"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestAccountBook(t *testing.T) {
	mockDb, mock := test.NewMockDB()
	defer test.CloseMockDB(mockDb)

	t.Run("Find 테스트", func(t *testing.T) {
		const findQuery = `SELECT \* FROM "account_book" WHERE id = \$1::uuid ORDER BY "account_book"."id" LIMIT \$2`

		userId := uuid.MustParse("11111111-1111-1111-1111-111111111111")
		accountBook, _ := accountBookModel.New(userId, "testAccountBook")
		role, _ := roleModel.New(userId, accountBook.Id, "owner")
		t.Run("AccountBook이 없다면 에러를 반환한다.", func(t *testing.T) {
			spec := domain_specs.NewDeletableAccountBookSpec([]*roleModel.Role{role}, accountBook.Id)

			mock.ExpectQuery(findQuery).
				WithArgs(accountBook.Id, 1).
				WillReturnError(gorm.ErrRecordNotFound)

			result, err := spec.Find(mockDb)
			assert.Nil(t, result)
			assert.Equal(t, fmt.Sprintf("Account book(%s) is not found.", accountBook.Id), err.Error())
		})

		t.Run("AccountBook이 이미 삭제되었다면 에러를 반환한다.", func(t *testing.T) {
			spec := domain_specs.NewDeletableAccountBookSpec([]*roleModel.Role{role}, accountBook.Id)

			deletedAt := time.Now()
			rows := sqlmock.NewRows([]string{"id", "deletedAt"}).
				AddRow(accountBook.Id, deletedAt)

			mock.ExpectQuery(findQuery).
				WithArgs(accountBook.Id, 1).
				WillReturnRows(rows)

			result, err := spec.Find(mockDb)
			assert.Nil(t, result)
			assert.Equal(t, fmt.Sprintf("Account book(%s) is already deleted.", accountBook.Id), err.Error())
		})

		t.Run("유저의 롤이 accountBook의 owner가 아니라면 에러를 반환한다.", func(t *testing.T) {
			memberRole, _ := roleModel.New(userId, accountBook.Id, "member")

			spec := domain_specs.NewDeletableAccountBookSpec([]*roleModel.Role{memberRole}, accountBook.Id)

			rows := sqlmock.NewRows([]string{"id"}).
				AddRow(accountBook.Id)

			mock.ExpectQuery(findQuery).
				WithArgs(accountBook.Id, 1).
				WillReturnRows(rows)

			result, err := spec.Find(mockDb)
			assert.Nil(t, result)
			assert.Equal(t, fmt.Sprintf("Actor has no permission to able to delete this account book(%s).", accountBook.Id), err.Error())
		})
	})
}
