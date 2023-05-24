package api

import (
	"errors"
	"opChat/entity"
	"opChat/errcode"
	"opChat/global"
	"opChat/request"
	"opChat/response"
	"sync/atomic"

	"gorm.io/gorm"
)

type message struct{}

var Message message

func (m *message) Send(uid int, request *request.MessageSend) (*response.Response[response.MessageSend], *errcode.Error) {
	tx := global.Database.Begin()
	var (
		err     error
		message entity.Message
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
		if group.Creator != uint(uid) {
			member := entity.Member{
				Group: request.ID,
				User:  uid,
				Grant: true,
			}
			err = tx.Where(&member, "grant").First(&member).Error
			if err != nil {
				tx.Rollback()
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return nil, errcode.UserNotInGroup
				}
				return nil, errcode.FindDataError.WithDetail(err.Error())
			}
		}
	} else {
		friend := entity.Friend{
			From:  uid,
			To:    request.ID,
			Grant: true,
		}
		err = tx.Where(&friend, "grant").First(&friend).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			friend := entity.Friend{
				From:  request.ID,
				To:    uid,
				Grant: true,
			}
			err = tx.Where(&friend, "grant").First(&friend).Error
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
	}
	message = entity.Message{
		ID:   uint(atomic.AddInt32(&global.NowMessageID, 1)),
		From: uid,
		To:   request.ID,
		Data: request.Data,
	}
	err = tx.Create(&message).Error
	if err != nil {
		tx.Rollback()
		return nil, errcode.InsertDataError.WithDetail(err.Error())
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return nil, errcode.CommitError.WithDetail(err.Error())
	}
	return &response.Response[response.MessageSend]{
		Code:    200,
		Message: "发送成功",
		Data: &response.MessageSend{
			ID: int(message.ID),
		},
	}, nil
}
