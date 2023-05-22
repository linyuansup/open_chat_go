package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID             uint `gorm:"primarykey"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
	PhoneNumber    string
	Username       string
	Password       string
	DeviceID       string
	AvatarFileName string
	AvatarExName   string
}
