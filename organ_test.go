package main

import (
	"context"
	"opChat/api"
	"opChat/request"
	"testing"
)

func TestOrganJoin(t *testing.T) {
	_, e := api.Organ.Join(userID2, &request.OrganJoin{
		ID: groupID,
	}, context.Background())
	if e != nil {
		t.Fatal("加入组织失败：" + e.Error())
	}
	_, e = api.Organ.Join(userID2, &request.OrganJoin{
		ID: groupID,
	}, context.Background())
	if e == nil {
		t.Fatal("第二次加入组织成功")
	}
	_, e = api.Organ.Join(userID2, &request.OrganJoin{
		ID: userID,
	}, context.Background())
	if e != nil {
		t.Fatal("加入组织失败：" + e.Error())
	}
	_, e = api.Organ.Join(userID2, &request.OrganJoin{
		ID: userID,
	}, context.Background())
	if e == nil {
		t.Fatal("第二次加入组织成功")
	}
}

func TestOrganAvatar(t *testing.T) {
	_, e := api.Organ.Avatar(userID, &request.OrganAvatar{
		ID: userID,
	}, context.Background())
	if e != nil {
		t.Fatal("获取头像失败：" + e.Error())
	}
	_, e = api.Organ.Avatar(userID, &request.OrganAvatar{
		ID: groupID,
	}, context.Background())
	if e != nil {
		t.Fatal("获取头像失败：" + e.Error())
	}
}

func TestOrganSetAvatar(t *testing.T) {
	_, e := api.Organ.SetAvatar(userID, &request.OrganSetAvatar{
		ID:   userID,
		File: avatarBase,
		Ex:   "png",
	}, context.Background())
	if e != nil {
		t.Fatal("设置头像失败：" + e.Error())
	}
	_, e = api.Organ.SetAvatar(userID, &request.OrganSetAvatar{
		ID:   groupID,
		File: avatarBase,
		Ex:   "png",
	}, context.Background())
	if e != nil {
		t.Fatal("设置头像失败：" + e.Error())
	}
}

func TestOrganName(t *testing.T) {
	res, e := api.Organ.Name(userID, &request.OrganName{
		ID: userID,
	}, context.Background())
	if e != nil {
		t.Fatal("获取用户名失败：" + e.Error())
	}
	if res.Data.Name != "新用户" {
		t.Fatal("用户名错误：" + res.Data.Name)
	}
	res, e = api.Organ.Name(userID, &request.OrganName{
		ID: groupID,
	}, context.Background())
	if e != nil {
		t.Fatal("获取用户名失败：" + e.Error())
	}
	if res.Data.Name != "测试群组" {
		t.Fatal("用户名错误：" + res.Data.Name)
	}
}

func TestOrganExit(t *testing.T) {
	_, e := api.Organ.Exit(userID2, &request.OrganExit{
		ID: groupID,
	}, context.Background())
	if e != nil {
		t.Fatal("退出组织失败：" + e.Error())
	}
}
