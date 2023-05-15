package main

import (
	"fmt"
	"opChat/global"
	"testing"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	phoneNumber = "19556172642"
	password = "81a0ad68ca3e7943db8833dc48927e2f"
	deviceID = "5wi1RhQ#JMunWu_I"
)

func TestMain(m *testing.M) {
	var err error
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", global.DatabaseAddress, global.DatabasePort, global.DatabaseUsername, global.DatabasePassword, global.DatabaseName+"_test")
	global.Database, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	initUserID()
	initDefaultAvatar()
	initDir()
	defer cleanUp()
	m.Run()
}

func cleanUp() {
	global.Database.Exec("delete from users")
}
