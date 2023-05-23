package api

import (
	"context"
	"errors"
	"opChat/entity"
	"opChat/errcode"
	"opChat/global"
	"opChat/request"
	"opChat/response"
	"opChat/util"

	"gorm.io/gorm"
)

type organ struct{}

var Organ organ

func (o *organ) Join(uid int, request *request.OrganJoinRequest, ctx context.Context) (*response.Response[response.OrganJoinResponse], *errcode.Error) {
	tx := global.Database.Begin()
	var err error
	if request.ID/100000000 >= 6 {
		targetGroup := entity.Group{
			ID: uint(request.ID),
		}
		err = tx.Find(&targetGroup).Error
		if err != nil {
			tx.Rollback()
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errcode.NoGroupFound
			}
			return nil, errcode.FindDataError.WithDetail(err.Error())
		}
		if targetGroup.Creator == uint(uid) {
			tx.Rollback()
			return nil, errcode.UserIsCreator
		}
		member := entity.Member{
			User:  uid,
			Group: request.ID,
		}
		err = tx.First(&member).Error
		if err == nil {
			tx.Rollback()
			if member.Admin {
				return nil, errcode.UserIsAdmin
			}
			if !member.Grant {
				return nil, errcode.AlreadyRequest
			}
			return nil, errcode.UserIsMember
		}
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return nil, errcode.FindDataError.WithDetail(err.Error())
		}
		member.Grant = false
		err = tx.Create(&member).Error
		if err != nil {
			tx.Rollback()
			return nil, errcode.CommitError.WithDetail(err.Error())
		}
	} else {
		friend := entity.Friend{
			From: uid,
			To:   request.ID,
		}
		err = tx.First(&friend).Error
		if err == nil {
			tx.Rollback()
			if !friend.Grant {
				return nil, errcode.AlreadyRequest
			}
			return nil, errcode.UserIsMember
		}
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return nil, errcode.FindDataError.WithDetail(err.Error())
		}
		friend = entity.Friend{
			From: request.ID,
			To:   uid,
		}
		err = tx.First(&friend).Error
		if err == nil {
			tx.Rollback()
			if !friend.Grant {
				return nil, errcode.AlreadyRequest
			}
			return nil, errcode.UserIsMember
		}
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return nil, errcode.FindDataError.WithDetail(err.Error())
		}
		friend.Grant = false
		err = tx.Create(&friend).Error
		if err != nil {
			return nil, errcode.InsertDataError.WithDetail(err.Error())
		}
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return nil, errcode.CommitError.WithDetail(err.Error())
	}
	return &response.Response[response.OrganJoinResponse]{
		Code:    200,
		Message: "申请成功",
		Data:    &response.OrganJoinResponse{},
	}, nil
}

func (o *organ) Avatar(uid int, request *request.OrganAvatarRequest, ctx context.Context) (*response.Response[response.OrganAvatarResponse], *errcode.Error) {
	tx := global.Database.Begin()
	var (
		avatarName string
		avatarEx   string
		err        error
	)
	if request.ID/100000000 >= 6 {
		group := entity.Group{
			ID: uint(request.ID),
		}
		err = tx.First(&group).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errcode.NoGroupFound
			}
			return nil, errcode.FindDataError.WithDetail(err.Error())
		}
		avatarName = group.AvatarFileName
		avatarEx = group.AvatarExName
	} else {
		user := entity.User{
			ID: uint(request.ID),
		}
		err = tx.First(&user).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errcode.NoUserRequestFound
			}
			return nil, errcode.FindDataError.WithDetail(err.Error())
		}
		avatarName = user.AvatarFileName
		avatarEx = user.AvatarExName
	}
	file, e := util.OpenFile(avatarName, avatarEx, "avatar")
	if e != nil {
		return nil, e
	}
	return &response.Response[response.OrganAvatarResponse]{
		Code:    200,
		Message: "获取成功",
		Data: &response.OrganAvatarResponse{
			File: util.Base64Encode(file),
			Ex:   avatarEx,
		},
	}, nil
}
