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
}
