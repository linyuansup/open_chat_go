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

type group struct{}

var Group group

func (g *group) Create(uid int, request *request.GroupCreateRequest, ctx context.Context) (*response.Response[response.GroupCreateResponse], *errcode.Error) {
	tx := global.Database.Begin()
	id := atomic.AddInt32(&global.NowGroupID, 1)
	e := tx.Create(&entity.Group{
		ID:             uint(id),
		Creator:        uint(uid),
		Name:           request.Name,
		AvatarFileName: global.AvatarFileName,
		AvatarExName:   global.AvatarExName,
	}).Error
	if e != nil {
		tx.Rollback()
		return nil, errcode.InsertDataError.WithDetail(e.Error())
	}
	err := tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return nil, errcode.CommitError.WithDetail(err.Error())
	}
	return &response.Response[response.GroupCreateResponse]{
		Code:    200,
		Message: "创建群聊成功",
		Data: &response.GroupCreateResponse{
			ID: int(id),
		},
	}, nil
}

func (g *group) Delete(uid int, request *request.GroupDeleteRequest, ctx context.Context) (*response.Response[response.GroupDeleteResponse], *errcode.Error) {
	tx := global.Database.Begin()
	targetGroup := entity.Group{
		ID: uint(request.ID),
	}
	err := tx.Find(&targetGroup).Error
	if err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errcode.NoGroupFound
		}
		return nil, errcode.FindDataError.WithDetail(err.Error())
	}
	if targetGroup.Creator != uint(uid) {
		tx.Rollback()
		return nil, errcode.NotCreator
	}
	err = tx.Where("members.group = ?", targetGroup.ID).Delete(&entity.Member{}).Error
	if err != nil {
		tx.Rollback()
		return nil, errcode.DeleteDataError.WithDetail(err.Error())
	}
	err = tx.Where("messages.to = ?", targetGroup.ID).Delete(&entity.Message{}).Error
	if err != nil {
		tx.Rollback()
		return nil, errcode.DeleteDataError.WithDetail(err.Error())
	}
	err = tx.Delete(&targetGroup).Error
	if err != nil {
		tx.Rollback()
		return nil, errcode.DeleteDataError.WithDetail(err.Error())
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return nil, errcode.CommitError.WithDetail(err.Error())
	}
	return &response.Response[response.GroupDeleteResponse]{
		Code:    200,
		Message: "删除群聊成功",
		Data:    &response.GroupDeleteResponse{},
	}, nil
}

func (g *group) Agree(uid int, request *request.GroupAgreeRequest, ctx context.Context) (*response.Response[response.GroupAgreeResponse], *errcode.Error) {
	tx := global.Database.Begin()
	targetGroup := entity.Group{
		ID: uint(request.GroupID),
	}
	err := tx.First(&targetGroup).Error
	if err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errcode.NoGroupFound
		}
		return nil, errcode.FindDataError.WithDetail(err.Error())
	}
	if uid != int(targetGroup.Creator) {
		member := entity.Member{
			Group: request.GroupID,
			User:  uid,
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
			return nil, errcode.NoChangePermission
		}
	}
	member := entity.Member{
		User:  request.UserID,
		Group: request.GroupID,
	}
	err = tx.First(&member).Error
	if err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errcode.NoRequest
		}
		return nil, errcode.FindDataError.WithDetail(err.Error())
	}
	if member.Grant {
		tx.Rollback()
		return nil, errcode.UserIsMember
	}
	member.Grant = true
	err = tx.Save(&member).Error
	if err != nil {
		tx.Rollback()
		return nil, errcode.UpdateDataError.WithDetail(err.Error())
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return nil, errcode.CommitError.WithDetail(err.Error())
	}
	return &response.Response[response.GroupAgreeResponse]{
		Code:    200,
		Message: "操作成功",
		Data:    &response.GroupAgreeResponse{},
	}, nil
}

func (g *group) SetAdmin(uid int, request *request.GroupSetAdminRequest, ctx context.Context) (*response.Response[response.GroupSetAdminResponse], *errcode.Error) {
	tx := global.Database.Begin()
	group := entity.Group {
		ID: uint(uid),
	}
	err := tx.Find(&group).Error
	if err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errcode.NoGroupFound
		}
		return nil, errcode.FindDataError.WithDetail(err.Error())
	}
	if group.Creator != uint(uid) {
		tx.Rollback()
		return nil, errcode.NoChangePermission
	}
	member := entity.Member {
		User: request.UserID,
		Group: request.GroupID,
	}
	err = tx.Find(&member).Error
	if err != nil{
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
	if !member.Admin {
		tx.Rollback()
		return nil, errcode.UserIsAdmin
	}
	member.Admin = true
	err = tx.Save(&member).Error
	if err != nil {
		tx.Rollback()
		return nil, errcode.UpdateDataError.WithDetail(err.Error())
	}
	err = tx.Commit().Error
	if err != nil {
		return nil, errcode.CommitError.WithDetail(err.Error())
	}
	return &response.Response[response.GroupSetAdminResponse]{
		Code:    200,
		Message: "设置成功",
		Data:    &response.GroupSetAdminResponse{},
	}, nil
}
