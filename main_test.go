package main

import (
	"opChat/global"
	"testing"
)

const (
	phoneNumber      = "19556172642"
	wrongPhoneNumber = "19556172643"
	password         = "81a0ad68ca3e7943db8833dc48927e2f"
	wrongPassword    = "81a0ad68ca3e7943db8833dc48927e2d"
	deviceID         = "5wi1RhQ#JMunWu_I"
	wrongDeviceID    = "5wi1RhQ#JMunWd_I"
)

var (
	userID  int
	groupID int
)

func TestMain(m *testing.M) {
	global.DatabaseAddress = "43.143.59.198"
	global.DatabaseName += "_test"
	global.Init()
	cleanUp()
	global.Init()
	m.Run()
}

func TestRunner(t *testing.T) {
	if !t.Run("TestUserCreate", TestUserCreate) ||
		!t.Run("TestUserLogin", TestUserLogin) ||
		!t.Run("TestUserSetPassword", TestUserSetPassword) ||
		!t.Run("TestGroupCreate", TestGroupCreate) ||
		!t.Run("TestGroupDelete", TestGroupDelete) {
		t.Fatal()
	}
}

func cleanUp() {
	global.Database.Exec("delete from users")
	global.Database.Exec("delete from groups")
	global.Database.Exec("delete from messages")
	global.Database.Exec("delete from relations")
}
