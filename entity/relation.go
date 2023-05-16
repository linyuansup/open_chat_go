package entity

import "gorm.io/gorm"

type Relation struct {
	gorm.Model
	SenderID   int
	RecieverID int
	Mode       int
}
