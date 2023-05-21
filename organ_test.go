package main

import (
	"context"
	"opChat/api"
	"opChat/request"
	"testing"
)

func TestOrganJoin(t *testing.T) {
	_, e := api.Organ.Join(userID, &request.OrganJoinRequest{
		ID: userID2,
	}, context.Background())
	if e != nil {
		t.Fatal("加好友失败，" + e.Error())
	}
	_, e = api.Organ.Join(userID, &request.OrganJoinRequest{
		ID: userID2,
	}, context.Background())
	if e == nil {
		t.Fatal("多次加好友成功")
	}
	_, e = api.Organ.Join(userID, &request.OrganJoinRequest{
		ID: groupID,
	}, context.Background())
	if e != nil {
		t.Fatal("加群失败，" + e.Error())
	}
	_, e = api.Organ.Join(userID, &request.OrganJoinRequest{
		ID: groupID,
	}, context.Background())
	if e == nil {
		t.Fatal("多次加群成功")
	}
}
