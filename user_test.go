package main

import (
	"context"
	"opChat/api"
	"opChat/request"
	"testing"
)

func TestUserCreate(t *testing.T) {
	res, e := api.User.Create(0, &request.UserCreate{
		PhoneNumber: phoneNumber,
		Password:    password,
		DeviceID:    deviceID,
	}, context.Background())
	if e != nil {
		t.Fatal("用户第一次注册失败：" + e.Error())
	}
	userID = int(res.Data.ID)
	_, e = api.User.Create(0, &request.UserCreate{
		PhoneNumber: phoneNumber,
		Password:    password,
		DeviceID:    deviceID,
	}, context.Background())
	if e == nil {
		t.Fatal("用户第二次注册成功")
	}
	res, e = api.User.Create(0, &request.UserCreate{
		PhoneNumber: phoneNumber2,
		Password:    password,
		DeviceID:    deviceID,
	}, context.Background())
	if e != nil {
		t.Fatal("用户第一次注册失败：" + e.Error())
	}
	userID2 = int(res.Data.ID)
}

func TestUserLogin(t *testing.T) {
	_, e := api.User.Login(0, &request.UserLogin{
		PhoneNumber: phoneNumber,
		Password:    password,
		DeviceID:    deviceID,
	}, context.Background())
	if e != nil {
		t.Fatal("用户登录失败：" + e.Error())
	}
	_, e = api.User.Login(0, &request.UserLogin{
		PhoneNumber: wrongPhoneNumber,
		Password:    password,
		DeviceID:    deviceID,
	}, context.Background())
	if e == nil {
		t.Fatal("用户使用错误手机号登录成功")
	}
	_, e = api.User.Login(0, &request.UserLogin{
		PhoneNumber: phoneNumber,
		Password:    wrongPassword,
		DeviceID:    deviceID,
	}, context.Background())
	if e == nil {
		t.Fatal("用户使用错误密码登录成功")
	}
	_, e = api.User.Login(0, &request.UserLogin{
		PhoneNumber: phoneNumber,
		Password:    password,
		DeviceID:    wrongDeviceID,
	}, context.Background())
	if e == nil {
		t.Fatal("用户使用错误设备码登录成功")
	}
}

func TestUserSetPassword(t *testing.T) {
	_, e := api.User.SetPassword(userID, &request.UserSetPassword{
		OldPassword: password,
		Password:    password,
	}, context.Background())
	if e != nil {
		t.Fatal("修改密码失败：" + e.Error())
	}
	_, e = api.User.SetPassword(userID, &request.UserSetPassword{
		OldPassword: wrongPassword,
		Password:    password,
	}, context.Background())
	if e == nil {
		t.Fatal("使用错误密码修改密码成功")
	}
}
