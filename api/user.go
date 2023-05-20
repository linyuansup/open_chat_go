package api

import (
	"context"
	"opChat/database"
	"opChat/entity"
	"opChat/errcode"
	"opChat/global"
	"opChat/request"
	"opChat/response"
	"sync/atomic"

	"gorm.io/gorm"
)

type user struct{}

var User user

func (u *user) Create(uid int, request *request.UserCreateRequest, ctx context.Context) (*response.Response[response.UserCreateResponse], *errcode.Error) {
	tx := global.Database.Begin()
	d := database.New[entity.User](tx, ctx)
	_, e := d.FindByField(&entity.User{PhoneNumber: request.PhoneNumber})
	if e == nil || e.Code != errcode.NoTargetFound.Code {
		tx.Rollback()
		return nil, errcode.PhoneNumberAlreadyExist
	}
	targetUser := &entity.User{
		Model: gorm.Model{
			ID: uint(atomic.AddInt32(&global.NowUserID, 1)),
		},
		PhoneNumber:    request.PhoneNumber,
		Password:       request.Password,
		DeviceID:       request.DeviceID,
		AvatarFileName: "e859977fae97b33c7e3e56d46098bd5d",
		AvatarExName:   "jpg",
	}
	e = d.Add(targetUser)
	if e != nil {
		tx.Rollback()
		return nil, e
	}
	err := tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return nil, errcode.CommitError.WithDetail(err.Error())
	}
	return &response.Response[response.UserCreateResponse]{
		Code:    200,
		Message: "注册成功",
		Data: &response.UserCreateResponse{
			ID: targetUser.ID,
		},
	}, nil
}

func (u *user) Login(uid int, request *request.UserLoginRequest, ctx context.Context) (*response.Response[response.UserLoginResponse], *errcode.Error) {
	targetUser, e := database.New[entity.User](global.Database, ctx).FindByField(&entity.User{PhoneNumber: request.PhoneNumber})
	if e != nil {
		if e.Code == errcode.NoTargetFound.Code {
			return nil, errcode.NoPhoneNumberFound
		}
		return nil, e
	}
	if (*targetUser)[0].Password != request.Password {
		return nil, errcode.WrongPassword
	}
	if (*targetUser)[0].DeviceID != request.DeviceID {
		return nil, errcode.WrongDeviceID
	}
	return &response.Response[response.UserLoginResponse]{
		Code:    200,
		Message: "登录成功",
		Data: &response.UserLoginResponse{
			ID: (*targetUser)[0].ID,
		},
	}, nil
}

func (u *user) SetPassword(uid int, request *request.UserSetPasswordRequest, ctx context.Context) (*response.Response[response.UserSetPasswordResponse], *errcode.Error) {
	tx := global.Database.Begin()
	d := database.New[entity.User](tx, ctx)
	targetUser, e := d.FindByID(uint(uid))
	if e != nil {
		if e.Code == errcode.NoUserFound.Code {
			tx.Rollback()
			return nil, errcode.NoTargetFound
		}
		tx.Rollback()
		return nil, e
	}
	if request.OldPassword != targetUser.Password {
		tx.Rollback()
		return nil, errcode.WrongPassword
	}
	e = d.Update(targetUser)
	if e != nil {
		tx.Rollback()
		return nil, e
	}
	err := tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return nil, errcode.CommitError.WithDetail(err.Error())
	}
	return &response.Response[response.UserSetPasswordResponse]{
		Code:    200,
		Message: "更改密码成功",
		Data:    &response.UserSetPasswordResponse{},
	}, nil
}
