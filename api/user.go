package api

import (
	"context"
	"opChat/database"
	"opChat/entity"
	"opChat/errcode"
	"opChat/request"
	"opChat/response"
)

type user struct{}

var User user

func (u *user) Create(uid int, request *request.UserCreateRequest, ctx context.Context) (*response.Response[response.UserCreateResponse], *errcode.Error) {
	_, e := database.UserDatabase.FindByPhoneNumber(request.PhoneNumber, ctx)
	if e == nil || e.Code != errcode.NoUserFound.Code {
		return nil, errcode.PhoneNumberAlreadyExist
	}
	targetUser := &entity.User{
		PhoneNumber: request.PhoneNumber,
		Password: request.Password,
		DeviceID: request.DeviceID,
	}
	e = database.UserDatabase.Add(targetUser, ctx)
	if e != nil {
		return nil, e
	}
	return &response.Response[response.UserCreateResponse]{
		Code: 200,
		Message: "注册成功",
		Data: &response.UserCreateResponse{
			ID: targetUser.ID,
		},
	}, nil
}

func (u *user) Login(uid int, request *request.UserLoginRequest, ctx context.Context) (*response.Response[response.UserLoginResponse], *errcode.Error) {
	targetUser, e := database.UserDatabase.FindByPhoneNumber(request.PhoneNumber,ctx)
	if e != nil {
		if e.Code == errcode.NoUserFound.Code {
			return nil, errcode.NoPhoneNumberFound
		}
		return nil, e
	}
	if targetUser.Password != request.Password {
		return nil, errcode.WrongPassword
	}
	if targetUser.DeviceID != request.DeviceID {
		return nil, errcode.WrongDeviceID
	}
	return &response.Response[response.UserLoginResponse]{
		Code: 200,
		Message: "登录成功",
		Data: &response.UserLoginResponse{
			ID: targetUser.ID,
		},
	}, nil
}
