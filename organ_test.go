package main

import (
	"context"
	"opChat/api"
	"opChat/request"
	"testing"
)

func TestOrganJoin(t *testing.T) {
	_, e := api.Oran.Join(userID2, &request.OrganJoinRequest{
		ID: groupID,
	}, context.Background())
	if e != nil {
		t.Fatal("加入组织失败：" + e.Error())
	}
	_, e = api.Oran.Join(userID2, &request.OrganJoinRequest{
		ID: groupID,
	}, context.Background())
	if e == nil {
		t.Fatal("第二次加入组织成功")
	}
	_, e = api.Oran.Join(userID2, &request.OrganJoinRequest{
		ID: userID,
	}, context.Background())
	if e != nil {
		t.Fatal("加入组织失败：" + e.Error())
	}
	_, e = api.Oran.Join(userID2, &request.OrganJoinRequest{
		ID: userID,
	}, context.Background())
	if e == nil {
		t.Fatal("第二次加入组织成功")
	}
}
