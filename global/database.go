package global

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func initDatabase() {
	var err error
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", DatabaseAddress, DatabasePort, DatabaseUsername, DatabasePassword, DatabaseName)
	Database, err = gorm.Open(postgres.Open(connStr), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic(err)
	}
	db, err := Database.DB()
	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(200)
	db.SetConnMaxLifetime(time.Minute)
	db.SetConnMaxIdleTime(time.Minute)
}
