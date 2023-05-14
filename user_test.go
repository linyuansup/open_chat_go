package main

import (
	"context"
	"opChat/api"
	"opChat/request"
	"testing"
)

func TestCreateUser(t *testing.T) {
	res, err := api.User.Create(0, &request.UserCreateRequest{
		PhoneNumber: "19556172642",
		Password: "81A0AD68CA3E7943DB8833DC48927E2F",
		DeviceID: "5wi1RhQ#JMunWu_I",
	}, context.Background())
	if err != nil {
		t.Logf("create user fail: %+v", err)
		t.FailNow()
	}
	t.Logf("create user success: %+v", *(res.Data))
}
