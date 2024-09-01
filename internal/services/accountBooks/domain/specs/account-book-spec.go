package domain_specs

import (
	"banana-account-book.com/internal/services/accountBooks/domain"
	"gorm.io/gorm"
)

type AccountBookSpec interface {
	Find(db *gorm.DB) ([]*domain.AccountBook, error)
	Count(db *gorm.DB) (int, error)
}
