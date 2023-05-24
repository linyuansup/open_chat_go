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
		From: request.ID,
		To:   uid,
	}
	err = tx.First(&friend).Error
	if err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errcode.NoRequest
		}
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
		From: request.ID,
		To:   uid,
	}
	err = tx.First(&friend).Error
	if err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errcode.NoRequest
		}
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

func (f *friend) Request(uid int, request *request.FriendRequest) (*response.Response[response.FriendRequest], *errcode.Error) {
	tx := global.Database.Begin()
	var (
		ids    []int
		friend []entity.Friend
	)
	err := tx.Where(&entity.Friend{
		To:    uid,
		Grant: false,
	}, "to", "grant").Find(&friend).Error
	if err != nil {
		tx.Rollback()
		return nil, errcode.FindDataError.WithDetail(err.Error())
	}
	for _, v := range friend {
		ids = append(ids, v.From)
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return nil, errcode.CommitError.WithDetail(err.Error())
	}
	return &response.Response[response.FriendRequest]{
		Code:    200,
		Message: "获取成功",
		Data: &response.FriendRequest{
			ID: ids,
		},
	}, nil
}
