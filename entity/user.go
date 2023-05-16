package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	PhoneNumber    string
	Username       string
	Password       string
	DeviceID       string
	AvatarFileName string
	AvatarExName   string
}
