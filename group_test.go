package main

import (
	"context"
	"opChat/api"
	"opChat/request"
	"testing"
)

func TestGroupCreate(t *testing.T) {
	_, e := api.Group.Create(userID, &request.GroupCreateRequest{
		Name: "测试群聊",
	}, context.Background())
	if e != nil {
		t.Fatal("创建群聊失败：" + e.Error())
	}
	group, e := api.Group.Create(userID, &request.GroupCreateRequest{
		Name: "测试群聊 2",
	}, context.Background())
	if e != nil {
		t.Fatal("创建群聊失败：" + e.Error())
	}
	groupID = group.Data.ID
}

func TestGroupDelete(t *testing.T) {
	_, e := api.Group.Delete(userID, &request.GroupDeleteRequest{
		ID: groupID,
	}, context.Background())
	if e != nil {
		t.Fatal("删除群聊失败：" + e.Error())
	}
	_, e = api.Group.Delete(userID, &request.GroupDeleteRequest{
		ID: groupID + 1,
	}, context.Background())
	if e == nil {
		t.Fatal("使用错误 ID 删除群聊成功")
	}
}
