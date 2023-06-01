package global

import (
	"fmt"

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
}
