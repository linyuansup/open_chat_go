package database

import (
	"context"

	"gorm.io/gorm"
)

type Database struct {
	DB  *gorm.DB
	Ctx context.Context
}
