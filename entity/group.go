package entity

import (
	"time"

	"gorm.io/gorm"
)

type Group struct {
	ID             uint `gorm:"primarykey"`
	CreatedAt      time.Time `gorm:"primaryKey"`
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
	Creator        uint
	Name           string
	AvatarFileName string
	AvatarExName   string
}
