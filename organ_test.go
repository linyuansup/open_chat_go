package main

import (
	"context"
	"opChat/api"
	"opChat/request"
	"testing"
)

func TestOrganJoin(t *testing.T) {
	_, e := api.Organ.Join(userID2, &request.OrganJoinRequest{
		ID: groupID,
	}, context.Background())
	if e != nil {
		t.Fatal("加入组织失败：" + e.Error())
	}
	_, e = api.Organ.Join(userID2, &request.OrganJoinRequest{
		ID: groupID,
	}, context.Background())
	if e == nil {
		t.Fatal("第二次加入组织成功")
	}
	_, e = api.Organ.Join(userID2, &request.OrganJoinRequest{
		ID: userID,
	}, context.Background())
	if e != nil {
		t.Fatal("加入组织失败：" + e.Error())
	}
	_, e = api.Organ.Join(userID2, &request.OrganJoinRequest{
		ID: userID,
	}, context.Background())
	if e == nil {
		t.Fatal("第二次加入组织成功")
	}
}

func TestOrganAvatar(t *testing.T) {
	_, e := api.Organ.Avatar(userID, &request.OrganAvatarRequest{
		ID: userID,
	}, context.Background())
	if e != nil {
		t.Fatal("获取头像失败：" + e.Error())
	}
	_, e = api.Organ.Avatar(userID, &request.OrganAvatarRequest{
		ID: groupID,
	}, context.Background())
	if e != nil {
		t.Fatal("获取头像失败：" + e.Error())
	}
}

func TestOrganSetAvatar(t *testing.T) {
	_, e := api.Organ.SetAvatar(userID, &request.OrganSetAvatarRequest{
		ID: userID,
		File: avatarBase,
		Ex: "png",
	}, context.Background())
	if e != nil {
		t.Fatal("设置头像失败：" + e.Error())
	}
	_, e = api.Organ.SetAvatar(userID, &request.OrganSetAvatarRequest{
		ID: groupID,
		File: avatarBase,
		Ex: "png",
	}, context.Background())
	if e != nil {
		t.Fatal("设置头像失败：" + e.Error())
	}
}

func TestOrganName(t *testing.T) {
	res, e := api.Organ.Name(userID, &request.OrganNameRequest{
		ID: userID,
	}, context.Background())
	if e != nil {
		t.Fatal("获取用户名失败：" + e.Error())
	}
	if res.Data.Name != "新用户" {
		t.Fatal("用户名错误：" + res.Data.Name)
	}
	res, e = api.Organ.Name(userID, &request.OrganNameRequest{
		ID: groupID,
	}, context.Background())
	if e != nil {
		t.Fatal("获取用户名失败：" + e.Error())
	}
	if res.Data.Name != "测试群组" {
		t.Fatal("用户名错误：" + res.Data.Name)
	}
}