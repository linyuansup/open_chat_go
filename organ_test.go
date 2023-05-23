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
	res, e := api.Organ.Avatar(userID, &request.OrganAvatarRequest{
		ID: userID,
	}, context.Background())
	if e != nil {
		t.Fatal("获取头像失败：" + e.Error())
	}
	t.Log(res.Data.File)
	res, e = api.Organ.Avatar(userID, &request.OrganAvatarRequest{
		ID: groupID,
	}, context.Background())
	if e != nil {
		t.Fatal("获取头像失败：" + e.Error())
	}
	t.Log(res.Data.File)
}
