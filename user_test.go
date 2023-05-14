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
		Password: "81a0ad68ca3e7943db8833dc48927e2f",
		DeviceID: "5wi1RhQ#JMunWu_I",
	}, context.Background())
	if err != nil {
		t.Logf("create user fail: %+v", err)
		t.FailNow()
	}
	t.Logf("create user success: %+v", *(res.Data))
}
