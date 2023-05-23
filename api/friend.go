package api

import (
	"errors"
	"opChat/entity"
	"opChat/errcode"
	"opChat/global"
	"opChat/request"
	"opChat/response"

	"gorm.io/gorm"
)

type friend struct{}

var Friend friend

func (f *friend) Agree(uid int, request *request.FriendAgree) (*response.Response[response.FriendAgree], *errcode.Error) {
	tx := global.Database.Begin()
	err := tx.First(&entity.User{
		ID: uint(request.ID),
	}).Error
	if err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errcode.NoUserRequestFound
		}
		return nil, errcode.FindDataError.WithDetail(err.Error())
	}
	friend := entity.Friend{
		From: uid,
		To:   request.ID,
	}
	err = tx.First(&friend).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		friend = entity.Friend{
			From: request.ID,
			To:   uid,
		}
		err = tx.First(&friend).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return nil, errcode.NoRequest
		}
		if err != nil {
			tx.Rollback()
			return nil, errcode.FindDataError.WithDetail(err.Error())
		}
	}
	if err != nil {
		tx.Rollback()
		return nil, errcode.FindDataError.WithDetail(err.Error())
	}
	if friend.Grant {
		tx.Rollback()
		return nil, errcode.UserIsFriend
	}
	friend.Grant = true
	err = tx.Save(&friend).Error
	if err != nil {
		tx.Rollback()
		return nil, errcode.UpdateDataError.WithDetail(err.Error())
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return nil, errcode.CommitError.WithDetail(err.Error())
	}
	return &response.Response[response.FriendAgree]{
		Code:    200,
		Message: "操作成功",
		Data:    &response.FriendAgree{},
	}, nil
}

func (f *friend) Disgree(uid int, request *request.FriendDisgree) (*response.Response[response.FriendDisgree], *errcode.Error) {
	tx := global.Database.Begin()
	err := tx.First(&entity.User{
		ID: uint(request.ID),
	}).Error
	if err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errcode.NoUserRequestFound
		}
		return nil, errcode.FindDataError.WithDetail(err.Error())
	}
	friend := entity.Friend{
		From: uid,
		To:   request.ID,
	}
	err = tx.First(&friend).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		friend = entity.Friend{
			From: request.ID,
			To:   uid,
		}
		err = tx.First(&friend).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return nil, errcode.NoRequest
		}
		if err != nil {
			tx.Rollback()
			return nil, errcode.FindDataError.WithDetail(err.Error())
		}
	}
	if err != nil {
		tx.Rollback()
		return nil, errcode.FindDataError.WithDetail(err.Error())
	}
	if friend.Grant {
		tx.Rollback()
		return nil, errcode.UserIsFriend
	}
	err = tx.Delete(&friend).Error
	if err != nil {
		tx.Rollback()
		return nil, errcode.UpdateDataError.WithDetail(err.Error())
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return nil, errcode.CommitError.WithDetail(err.Error())
	}
	return &response.Response[response.FriendDisgree]{
		Code:    200,
		Message: "操作成功",
		Data:    &response.FriendDisgree{},
	}, nil
}
