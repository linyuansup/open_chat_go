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
