package api

import (
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

func (o *organ) Join(uid int, request *request.OrganJoin) (*response.Response[response.OrganJoinResponse], *errcode.Error) {
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
			From:  request.ID,
			To:    uid,
			Grant: true,
		}
		err = tx.Where(&friend, "grant").First(&friend).Error
		if err == nil {
			tx.Rollback()
			return nil, errcode.UserIsFriend
		}
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return nil, errcode.FindDataError.WithDetail(err.Error())
		}
		friend = entity.Friend{
			From: uid,
			To:   request.ID,
		}
		err = tx.First(&friend).Error
		if err == nil {
			tx.Rollback()
			if friend.Grant {
				return nil, errcode.UserIsFriend
			}
			return nil, errcode.AlreadyRequest
		}
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return nil, errcode.FindDataError.WithDetail(err.Error())
		}
		friend.Grant = false
		err = tx.Create(&friend).Error
		if err != nil {
			tx.Rollback()
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

func (o *organ) Avatar(uid int, request *request.OrganAvatar) (*response.Response[response.OrganAvatar], *errcode.Error) {
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
	return &response.Response[response.OrganAvatar]{
		Code:    200,
		Message: "获取成功",
		Data: &response.OrganAvatar{
			File: util.Base64Encode(file),
			Ex:   avatarEx,
		},
	}, nil
}

func (o *organ) SetAvatar(uid int, request *request.OrganSetAvatar) (*response.Response[response.OrganSetAvatar], *errcode.Error) {
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
	return &response.Response[response.OrganSetAvatar]{
		Code:    200,
		Message: "更改头像成功",
		Data: &response.OrganSetAvatar{
			Name: name,
		},
	}, nil
}

func (o *organ) Name(uid int, request *request.OrganName) (*response.Response[response.OrganName], *errcode.Error) {
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
	return &response.Response[response.OrganName]{
		Code:    200,
		Message: "获取成功",
		Data: &response.OrganName{
			Name: name,
		},
	}, nil
}

func (o *organ) Exit(uid int, request *request.OrganExit) (*response.Response[response.OrganExit], *errcode.Error) {
	tx := global.Database.Begin()
	var err error
	if request.ID >= 600000000 {
		group := entity.Group{
			ID: uint(request.ID),
		}
		err := tx.First(&group).Error
		if err != nil {
			tx.Rollback()
			return nil, errcode.NoGroupFound
		}
		if group.Creator == uint(uid) {
			tx.Rollback()
			return nil, errcode.UserIsCreator
		}
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
		if !member.Grant {
			tx.Rollback()
			return nil, errcode.UserNotInGroup
		}
		err = tx.Delete(&member).Error
		if err != nil {
			tx.Rollback()
			return nil, errcode.DeleteDataError.WithDetail(err.Error())
		}
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
				return nil, errcode.UserNotFriend
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
		if !friend.Grant {
			tx.Rollback()
			return nil, errcode.UserNotFriend
		}
		err = tx.Delete(&friend).Error
		if err != nil {
			tx.Rollback()
			return nil, errcode.DeleteDataError.WithDetail(err.Error())
		}
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return nil, errcode.CommitError.WithDetail(err.Error())
	}
	return &response.Response[response.OrganExit]{
		Code:    200,
		Message: "退出成功",
		Data:    &response.OrganExit{},
	}, nil
}
