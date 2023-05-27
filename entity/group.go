package entity

import (
	"gorm.io/gorm"
)

type Group struct {
	gorm.Model
	Creator        uint
	Name           string
	AvatarFileName string
	AvatarExName   string
}
