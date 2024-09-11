package domain_specs

import (
	"fmt"

	appError "banana-account-book.com/internal/libs/app-error"
	httpCode "banana-account-book.com/internal/libs/http/code"
	"banana-account-book.com/internal/services/accountBooks/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReadableAccountBookSpec struct {
	accountBookId uuid.UUID
}

func NewReadableAccountBookSpec(accountBookId uuid.UUID) *ReadableAccountBookSpec {
	return &ReadableAccountBookSpec{
		accountBookId: accountBookId,
	}
}

func (spec *ReadableAccountBookSpec) Find(db *gorm.DB) ([]*domain.AccountBook, error) {
	var accountBook *domain.AccountBook
	if err := db.Unscoped().Where("id = ?::uuid", spec.accountBookId).First(&accountBook).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, appError.New(httpCode.NotFound, fmt.Sprintf("Account book(%s) is not found.", spec.accountBookId), "가계부를 찾을 수 없습니다.")
		}

		return nil, appError.Wrap(err)
	}

	if accountBook.DeletedAt != nil {
		return nil, appError.New(httpCode.NotFound, fmt.Sprintf("Account book(%s) is already deleted.", spec.accountBookId), "가계부가 이미 삭제되었습니다.")
	}

	return []*domain.AccountBook{accountBook}, nil
}

func (spec *ReadableAccountBookSpec) Count(db *gorm.DB) (int, error) {
	return 0, appError.New(httpCode.NotImplemented, "Count method is not implemented", "")
}
