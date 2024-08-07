package entity

import (
	"time"

	"gorm.io/gorm"
)

type Aggregate struct {
	CreatedAt time.Time `json:"_", gorm:"autoCreateTime:nano;"`
	UpdatedAt time.Time `json:"__", gorm:"autoUpdateTime:nano;"`
}

type SoftDeletableAggregate struct {
	Aggregate
	DeletedAt gorm.DeletedAt `json:"_", gorm:"index;"`
}
