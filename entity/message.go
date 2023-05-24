package entity

import (
	"time"

	"gorm.io/gorm"
)

type Message struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time `gorm:"primaryKey"`
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	From      int
	To        int
	Data      string
}
