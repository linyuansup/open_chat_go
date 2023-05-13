package global

import (
	"fmt"
	"opChat/errcode"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func initDatabase() *errcode.Error {
	var err error
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", databaseAddress, databasePort, databaseUsername, databasePassword, databaseName)
	Database, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		return errcode.DatabaseConnectFail.WithDetail(err.Error())
	}
	return nil
}
