package domain_specs

import (
	"fmt"

	appError "banana-account-book.com/internal/libs/app-error"
	httpCode "banana-account-book.com/internal/libs/http/code"
	"banana-account-book.com/internal/services/accountBooks/domain"
	roleModel "banana-account-book.com/internal/services/roles/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DeletableAccountBookSpec struct {
	roles []*roleModel.Role

	accountBookId uuid.UUID
}

func NewDeletableAccountBookSpec(roles []*roleModel.Role, accountBookId uuid.UUID) *DeletableAccountBookSpec {
	return &DeletableAccountBookSpec{
		roles:         roles,
		accountBookId: accountBookId,
	}
}

func (spec *DeletableAccountBookSpec) Find(db *gorm.DB) ([]*domain.AccountBook, error) {
	var accountBook *domain.AccountBook
	if err := db.Unscoped().Where("id = ?::uuid", spec.accountBookId).First(&accountBook).Error; err != nil {
		return nil, appError.Wrap(err)
	}

	if accountBook.DeletedAt != nil {
		return nil, appError.New(httpCode.NotFound, fmt.Sprintf("Account book(%s) is already deleted", spec.accountBookId), "가계부가 이미 삭제되었습니다.")
	}

	if spec.IsOwner(accountBook) {
		return []*domain.AccountBook{accountBook}, nil
	}

	return nil, appError.New(httpCode.Forbidden, fmt.Sprintf("Actor has no permission to able to delete this account book(%s)", accountBook.Id), "가계부를 삭제할 권한이 없습니다.")
}

func (spec *DeletableAccountBookSpec) Count(db *gorm.DB) (int, error) {
	return 0, appError.New(httpCode.NotImplemented, "Count method is not implemented", "")
}

func (spec *DeletableAccountBookSpec) IsOwner(accountBook *domain.AccountBook) bool {
	for _, role := range spec.roles {
		if role.Type == "owner" && role.AccountBookId == accountBook.Id {
			return true
		}
	}
	return false
}
