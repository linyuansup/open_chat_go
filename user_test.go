package main

import (
	"context"
	"opChat/api"
	"opChat/request"
	"testing"
)

func TestCreateUser(t *testing.T) {
	res, err := api.User.Create(0, &request.UserCreateRequest{
		PhoneNumber: phoneNumber,
		Password: password,
		DeviceID: deviceID,
	}, context.Background())
	if err != nil {
		t.Logf("create user fail: %+v", err)
		t.FailNow()
	}
	t.Logf("create user success: %+v", *(res.Data))
}
