package entity

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	Data       string
	RelationID int
}
