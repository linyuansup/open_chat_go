package entity

import (
	"time"

	"gorm.io/gorm"
)

type Member struct {
	User      int
	Group     int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Grant     bool
	Admin     bool
}
