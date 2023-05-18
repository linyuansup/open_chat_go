package entity

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	Name           string
	AvatarFileName string
	AvatarExName   string
}
