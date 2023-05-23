package api

import (
	"context"
	"errors"
	"opChat/entity"
	"opChat/errcode"
	"opChat/global"
	"opChat/request"
	"opChat/response"

	"gorm.io/gorm"
)

type organ struct{}

var Oran organ

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
