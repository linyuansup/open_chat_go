package main

import (
	"opChat/global"
	"os"
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
	avatarBase       = "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"
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
		!t.Run("TestOrganAvatar", TestOrganAvatar) ||
		!t.Run("TestOrganSetAvatar", TestOrganSetAvatar) ||
		!t.Run("TestUserSetName", TestUserSetName) ||
		!t.Run("TestOrganExit", TestOrganExit) {
		t.Fatal()
	}
}

func cleanUp() {
	global.Database.Exec("delete from members").
		Exec("delete from groups").
		Exec("delete from messages").
		Exec("delete from friends").
		Exec("delete from users")
	os.RemoveAll("." + global.FilePath + "/avatar/c68c3bdebaf6d78d653de4e6e8c62e3c.png")
}
