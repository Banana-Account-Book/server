package infrastructure

import (
	"fmt"

	appError "banana-account-book.com/internal/libs/app-error"
	httpCode "banana-account-book.com/internal/libs/http/code"
	"banana-account-book.com/internal/services/roles/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RoleRepository interface {
	Save(db *gorm.DB, Role *domain.Role) error
	FindByUserId(db *gorm.DB, userId uuid.UUID) ([]*domain.Role, bool, error)
}

type RoleRepositoryImpl struct {
	manager *gorm.DB
}

func NewRoleRepository(manager *gorm.DB) RoleRepository {
	return &RoleRepositoryImpl{manager: manager}
}

func (r *RoleRepositoryImpl) Save(db *gorm.DB, role *domain.Role) error {
	if db == nil {
		db = r.manager
	}

	if err := db.Save(role).Error; err != nil {
		return appError.New(httpCode.InternalServerError, fmt.Sprintf("Failed to save role. %s", err.Error()), "")
	}
	return nil
}

func (r *RoleRepositoryImpl) FindByUserId(db *gorm.DB, userId uuid.UUID) ([]*domain.Role, bool, error) {
	if db == nil {
		db = r.manager
	}

	roles := []*domain.Role{}
	if err := db.Where("\"userId\" = ?::uuid", userId).Find(&roles).Error; err != nil {
		return nil, false, appError.New(httpCode.InternalServerError, fmt.Sprintf("Failed to findByUserId role. %s", err.Error()), "")
	}
	if len(roles) == 0 {
		return nil, false, nil
	}
	return roles, true, nil
}
