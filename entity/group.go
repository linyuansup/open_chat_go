package entity

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	Name           string
	CreatorID      int
	AvatarFileName string
	AvatarExName   string
}
