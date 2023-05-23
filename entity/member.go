package entity

import (
	"time"

	"gorm.io/gorm"
)

type Member struct {
	User      int `gorm:"primarykey"`
	Group     int `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Grant     bool
	Admin     bool
}
