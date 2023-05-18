package main

import (
	"context"
	"opChat/api"
	"opChat/request"
	"testing"
)

func TestOrganGetAvatar(t *testing.T) {
	_, e := api.Organ.GetAvatar(userID, &request.GetAvatarRequest{
		ID: userID,
	}, context.Background())
	if e != nil {
		t.Fatal("获取头像失败：" + e.Error())
	}
	_, e = api.Organ.GetAvatar(userID, &request.GetAvatarRequest{
		ID: userID - 1,
	}, context.Background())
	if e == nil {
		t.Fatal("使用错误 ID 获取头像成功")
	}
}
