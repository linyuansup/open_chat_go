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
	if request.ID >= 600000000 {
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
	if request.ID >= 600000000 {
		group := entity.Group{
			ID: uint(request.ID),
		}
		err = tx.First(&group).Error
		if err != nil {
			tx.Rollback()
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
			tx.Rollback()
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
		tx.Rollback()
		return nil, e
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return nil, errcode.CommitError.WithDetail(err.Error())
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

func (o *organ) SetAvatar(uid int, request *request.OrganSetAvatarRequest, ctx context.Context) (*response.Response[response.OrganSetAvatarResponse], *errcode.Error) {
	tx := global.Database.Begin()
	var err error
	file, e := util.Base64Decode([]byte(request.File))
	if e != nil {
		tx.Rollback()
		return nil, e
	}
	name, e := util.SaveFile(file, request.Ex, "avatar")
	if e != nil {
		tx.Rollback()
		return nil, e
	}
	if request.ID >= 600000000 {
		group := entity.Group{
			ID: uint(request.ID),
		}
		err = tx.First(&group).Error
		if err != nil {
			tx.Rollback()
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errcode.NoGroupFound
			}
			return nil, errcode.FindDataError.WithDetail(err.Error())
		}
		if group.Creator != uint(uid) {
			member := entity.Member{
				User:  uid,
				Group: request.ID,
			}
			err = tx.First(&member).Error
			if err != nil {
				tx.Rollback()
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return nil, errcode.UserNotInGroup
				}
				return nil, errcode.FindDataError.WithDetail(err.Error())
			}
			if !member.Admin {
				tx.Rollback()
				return nil, errcode.UserIsNotAdmin
			}
		}
		group.AvatarFileName = name
		group.AvatarExName = request.Ex
		err = tx.Save(&group).Error
		if err != nil {
			tx.Rollback()
			return nil, errcode.UpdateDataError.WithDetail(err.Error())
		}
	} else {
		if uid != request.ID {
			tx.Rollback()
			return nil, errcode.NoChangePermission
		}
		user := entity.User{
			ID: uint(uid),
		}
		err := tx.First(&user).Error
		if err != nil {
			tx.Rollback()
			return nil, errcode.FindDataError.WithDetail(err.Error())
		}
		user.AvatarFileName = name
		user.AvatarExName = request.Ex
		err = tx.Save(&user).Error
		if err != nil {
			tx.Rollback()
			return nil, errcode.UpdateDataError.WithDetail(err.Error())
		}
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return nil, errcode.CommitError.WithDetail(err.Error())
	}
	return &response.Response[response.OrganSetAvatarResponse]{
		Code:    200,
		Message: "更改头像成功",
		Data: &response.OrganSetAvatarResponse{
			Name: name,
		},
	}, nil
}

func (o *organ) Name(uid int, request *request.OrganNameRequest, ctx context.Context) (*response.Response[response.OrganNameResponse], *errcode.Error) {
	tx := global.Database.Begin()
	var (
		name string
		err  error
	)
	if request.ID >= 600000000 {
		group := entity.Group{
			ID: uint(request.ID),
		}
		err = tx.First(&group).Error
		if err != nil {
			tx.Rollback()
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errcode.NoGroupFound
			}
			return nil, errcode.FindDataError.WithDetail(err.Error())
		}
		name = group.Name
	} else {
		user := entity.User{
			ID: uint(request.ID),
		}
		err = tx.First(&user).Error
		if err != nil {
			tx.Rollback()
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errcode.NoUserRequestFound
			}
			return nil, errcode.FindDataError.WithDetail(err.Error())
		}
		name = user.Username
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return nil, errcode.CommitError.WithDetail(err.Error())
	}
	return &response.Response[response.OrganNameResponse]{
		Code:    200,
		Message: "获取成功",
		Data: &response.OrganNameResponse{
			Name: name,
		},
	}, nil
}
