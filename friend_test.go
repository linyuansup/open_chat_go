package main

import (
	"opChat/api"
	"opChat/request"
	"testing"
)

func TestFriendAgree(t *testing.T) {
	_, e := api.Friend.Agree(userID, &request.FriendAgree{
		ID: userID2,
	})
	if e != nil {
		t.Fatal("同意失败：" + e.Error())
	}
}

func TestFriendDisgree(t *testing.T) {
	_, e := api.Friend.Disgree(userID, &request.FriendDisgree{
		ID: userID3,
	})
	if e != nil {
		t.Fatal("不同意失败：" + e.Error())
	}
}
