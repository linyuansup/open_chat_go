package main

import (
	"context"
	"opChat/api"
	"opChat/request"
	"testing"
)

func TestUserCreate(t *testing.T) {
	_, e := api.User.Create(0, &request.UserCreateRequest{
		PhoneNumber: phoneNumber,
		Password:    password,
		DeviceID:    deviceID,
	}, context.Background())
	if e != nil {
		t.Fatal("用户第一次注册失败：" + e.Error())
	}
	_, e = api.User.Create(0, &request.UserCreateRequest{
		PhoneNumber: phoneNumber,
		Password:    password,
		DeviceID:    deviceID,
	}, context.Background())
	if e == nil {
		t.Fatal("用户第二次注册成功")
	}
}

func TestUserLogin(t *testing.T) {
	_, e := api.User.Login(0, &request.UserLoginRequest{
		PhoneNumber: phoneNumber,
		Password:    password,
		DeviceID:    deviceID,
	}, context.Background())
	if e != nil {
		t.Fatal("用户登录失败：" + e.Error())
	}
	_, e = api.User.Login(0, &request.UserLoginRequest{
		PhoneNumber: wrongPhoneNumber,
		Password:    password,
		DeviceID:    deviceID,
	}, context.Background())
	if e == nil {
		t.Fatal("用户使用错误手机号登录成功")
	}
	_, e = api.User.Login(0, &request.UserLoginRequest{
		PhoneNumber: phoneNumber,
		Password:    wrongPassword,
		DeviceID:    deviceID,
	}, context.Background())
	if e == nil {
		t.Fatal("用户使用错误密码登录成功")
	}
	_, e = api.User.Login(0, &request.UserLoginRequest{
		PhoneNumber: phoneNumber,
		Password:    password,
		DeviceID:    wrongDeviceID,
	}, context.Background())
	if e == nil {
		t.Fatal("用户使用错误设备码登录成功")
	}
}
