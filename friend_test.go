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

func TestFriendRequest(t *testing.T) {
	_, e := api.Organ.Join(userID3, &request.OrganJoin{
		ID: userID,
	})
	if e != nil {
		t.Fatal("申请失败：" + e.Error())
	}
	res, e := api.Friend.Request(userID, &request.FriendRequest{})
	if e != nil {
		t.Fatal("获取失败：" + e.Error())
	}
	t.Logf("%+v", res.Data.ID)
}
