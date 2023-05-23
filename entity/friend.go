package entity

import (
	"time"

	"gorm.io/gorm"
)

type Friend struct {
	From      int `gorm:"primarykey"`
	To        int `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Grant     bool
}