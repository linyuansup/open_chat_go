package entity

import (
	"time"

	"gorm.io/gorm"
)

type Member struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	User      int            `gorm:"primarykey"`
	Group     int            `gorm:"primarykey"`
	Grant     bool
	Admin     bool
}
