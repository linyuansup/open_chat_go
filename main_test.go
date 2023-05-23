package main

import (
	"opChat/global"
	"testing"
)

const (
	phoneNumber      = "19556172642"
	phoneNumber2     = "19556172644"
	wrongPhoneNumber = "19556172643"
	password         = "81a0ad68ca3e7943db8833dc48927e2f"
	wrongPassword    = "81a0ad68ca3e7943db8833dc48927e2d"
	deviceID         = "5wi1RhQ#JMunWu_I"
	wrongDeviceID    = "5wi1RhQ#JMunWd_I"
)

var (
	userID   int
	userID2  int
	groupID  int
	groupID2 int
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
		!t.Run("TestGroupDelete", TestGroupDelete) ||
		!t.Run("TestOrganJoin", TestOrganJoin) ||
		!t.Run("TestGroupAgree", TestGroupAgree) ||
		!t.Run("TestGroupSetAdmin", TestGroupSetAdmin) ||
		!t.Run("TestGroupRemoveAdmin", TestGroupRemoveAdmin) ||
		!t.Run("TestOrganAvatar", TestOrganAvatar) {
		t.Fatal()
	}
}

func cleanUp() {
	global.Database.Exec("delete from members").
		Exec("delete from groups").
		Exec("delete from messages").
		Exec("delete from friends").
		Exec("delete from users")
}
