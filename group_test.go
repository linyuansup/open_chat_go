package main

import (
	"context"
	"opChat/api"
	"opChat/request"
	"testing"
)

func TestGroupCreate(t *testing.T) {
	res, e := api.Group.Create(userID, &request.GroupCreateRequest{
		Name: "测试群组",
	}, context.Background())
	if e != nil {
		t.Fatal("创建群聊失败：" + e.Error())
	}
	groupID = res.Data.ID
	res, e = api.Group.Create(userID, &request.GroupCreateRequest{
		Name: "测试群组 2",
	}, context.Background())
	if e != nil {
		t.Fatal("创建群聊失败：" + e.Error())
	}
	groupID2 = res.Data.ID
}

func TestGroupDelete(t *testing.T) {
	_, e := api.Group.Delete(userID2, &request.GroupDeleteRequest{
		ID: groupID2,
	}, context.Background())
	if e == nil {
		t.Fatal("使用非群主删除群聊成功")
	}
	_, e = api.Group.Delete(userID, &request.GroupDeleteRequest{
		ID: groupID,
	}, context.Background())
	if e != nil {
		t.Fatal("删除群聊失败：" + e.Error())
	}
}
