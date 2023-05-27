package entity

import (
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	From int
	To   int
	Data string
}
