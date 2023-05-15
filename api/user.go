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
	if e.Code != errcode.NoUserFound.Code {
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
