package ddd

import (
	"time"

	"gorm.io/gorm"
)

type Aggregate struct {
	CreatedAt time.Time `json:"-" gorm:"column:createdAt;autoCreateTime:nano"`
	UpdatedAt time.Time `json:"-" gorm:"column:updatedAt;autoUpdateTime:nano"`
}

type SoftDeletableAggregate struct {
	Aggregate
	DeletedAt *gorm.DeletedAt `json:"-" gorm:"column:deletedAt"`
}
