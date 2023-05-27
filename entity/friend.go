package entity

import (
	"time"

	"gorm.io/gorm"
)

type Friend struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	From  int `gorm:"primarykey"`
	To    int `gorm:"primarykey"`
	Grant bool
}
