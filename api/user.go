package api

import (
	"context"
	"errors"
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
	err := tx.Where("phone_number = ?", request.PhoneNumber).First(&entity.User{}).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		return nil, errcode.PhoneNumberAlreadyExist
	}
	targetUser := &entity.User{
		ID:             uint(atomic.AddInt32(&global.NowUserID, 1)),
		Username:       "新用户",
		PhoneNumber:    request.PhoneNumber,
		Password:       request.Password,
		DeviceID:       request.DeviceID,
		AvatarFileName: global.AvatarFileName,
		AvatarExName:   global.AvatarExName,
	}
	err = tx.Create(targetUser).Error
	if err != nil {
		tx.Rollback()
		return nil, errcode.InsertDataError.WithDetail(err.Error())
	}
	err = tx.Commit().Error
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
	var targetUser entity.User
	e := global.Database.Where("phone_number = ?", request.PhoneNumber).Find(&targetUser)
	if e.RowsAffected == 0 {
		return nil, errcode.NoPhoneNumberFound
	}
	if e.Error != nil {
		return nil, errcode.FindDataError.WithDetail(e.Error.Error())
	}
	if targetUser.Password != request.Password {
		return nil, errcode.WrongPassword
	}
	if targetUser.DeviceID != request.DeviceID {
		return nil, errcode.WrongDeviceID
	}
	return &response.Response[response.UserLoginResponse]{
		Code:    200,
		Message: "登录成功",
		Data: &response.UserLoginResponse{
			ID: targetUser.ID,
		},
	}, nil
}

func (u *user) SetPassword(uid int, request *request.UserSetPasswordRequest, ctx context.Context) (*response.Response[response.UserSetPasswordResponse], *errcode.Error) {
	tx := global.Database.Begin()
	targetUser := entity.User{
		ID: uint(uid),
	}
	err := tx.First(&targetUser).Error
	if err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errcode.NoTargetFound
		}
		return nil, errcode.FindDataError.WithDetail(err.Error())
	}
	if request.OldPassword != targetUser.Password {
		tx.Rollback()
		return nil, errcode.WrongPassword
	}
	err = tx.Save(&targetUser).Error
	if err != nil {
		tx.Rollback()
		return nil, errcode.UpdateDataError.WithDetail(err.Error())
	}
	err = tx.Commit().Error
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
