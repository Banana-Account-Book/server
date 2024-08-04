package entity

import "time"

type Aggregate struct {
	CreatedAt time.Time `json:"createdAt" gorm:"column:createdAt"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updatedAt"`
}

type SoftDeletableAggregate struct {
	Aggregate
	DeletedAt *time.Time `json:"deletedAt" gorm:"column:deletedAt"`
}
