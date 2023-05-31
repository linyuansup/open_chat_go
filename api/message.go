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
			Model: gorm.Model{
				ID: uint(request.ID),
			},
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
		Model: gorm.Model{
			ID: uint(atomic.AddInt32(&global.NowMessageID, 1)),
		},
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

func (m *message) Up(uid int, request *request.MessageUp) (*response.Response[response.MessageUp], *errcode.Error) {
	tx := global.Database.Begin()
	var (
		msg []entity.Message
		err error
	)
	if request.ID >= 600000000 {
		group := entity.Group{
			Model: gorm.Model{
				ID: uint(request.ID),
			},
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
		if request.MsgID == 0 {
			err = tx.Where("messages.to = ?", request.ID).Order("messages.created_at DESC").Limit(request.Num).Find(&msg).Error
			if err != nil {
				tx.Rollback()
				return nil, errcode.FindDataError.WithDetail(err.Error())
			}
		} else {
			targetMsg := entity.Message{
				Model: gorm.Model{
					ID: uint(request.MsgID),
				},
			}
			err = tx.First(&targetMsg).Error
			if err != nil {
				tx.Rollback()
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return nil, errcode.NoMessageFound
				}
				return nil, errcode.FindDataError.WithDetail(err.Error())
			}
			if targetMsg.To != request.ID {
				tx.Rollback()
				return nil, errcode.MessageNotBelongGroup
			}
			err = tx.Where("messages.to = ? AND messages.created_at < ?", request.ID, targetMsg.CreatedAt).Order("messages.created_at DESC").Limit(request.Num).Find(&msg).Error
			if err != nil {
				tx.Rollback()
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
		if request.MsgID == 0 {
			err = tx.Where("(messages.from = ? AND messages.to = ?) OR (messages.from = ? AND messages.to = ?)", uid, request.ID, request.ID, uid).Order("messages.created_at DESC").Limit(request.Num).Find(&msg).Error
			if err != nil {
				tx.Rollback()
				return nil, errcode.FindDataError.WithDetail(err.Error())
			}
		} else {
			targetMsg := entity.Message{
				Model: gorm.Model{
					ID: uint(request.MsgID),
				},
			}
			err = tx.First(&targetMsg).Error
			if err != nil {
				tx.Rollback()
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return nil, errcode.NoMessageFound
				}
				return nil, errcode.FindDataError.WithDetail(err.Error())
			}
			if (targetMsg.From != uid && targetMsg.To != request.ID && targetMsg.From != request.ID && targetMsg.To != uid) || uid == request.ID {
				tx.Rollback()
				return nil, errcode.MessageNotBelongGroup
			}
			err = tx.Where("((messages.from = ? AND messages.to = ?) OR (messages.from = ? AND messages.to = ?)) AND messages.created_at < ?", uid, request.ID, request.ID, uid, targetMsg.CreatedAt).Order("messages.created_at DESC").Limit(request.Num).Find(&msg).Error
			if err != nil {
				tx.Rollback()
				return nil, errcode.FindDataError.WithDetail(err.Error())
			}
		}
	}
	result := []response.Message{}
	for _, v := range msg {
		result = append(result, response.Message{
			ID:     int(v.ID),
			Data:   v.Data,
			Sender: v.From,
			Time:   int(v.CreatedAt.Unix()),
		})
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return nil, errcode.CommitError.WithDetail(err.Error())
	}
	return &response.Response[response.MessageUp]{
		Code:    200,
		Message: "获取成功",
		Data: &response.MessageUp{
			Msg: result,
		},
	}, nil
}

func (m *message) Down(uid int, request *request.MessageDown) (*response.Response[response.MessageDown], *errcode.Error) {
	tx := global.Database.Begin()
	var (
		msg []entity.Message
		err error
	)
	if request.ID >= 600000000 {
		group := entity.Group{
			Model: gorm.Model{
				ID: uint(request.ID),
			},
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
		if request.MsgID == 0 {
			err = tx.Where("messages.to = ?", request.ID).Order("messages.created_at ASC").Limit(request.Num).Find(&msg).Error
			if err != nil {
				tx.Rollback()
				return nil, errcode.FindDataError.WithDetail(err.Error())
			}
		} else {
			targetMsg := entity.Message{
				Model: gorm.Model{
					ID: uint(request.MsgID),
				},
			}
			err = tx.First(&targetMsg).Error
			if err != nil {
				tx.Rollback()
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return nil, errcode.NoMessageFound
				}
				return nil, errcode.FindDataError.WithDetail(err.Error())
			}
			if targetMsg.To != request.ID {
				tx.Rollback()
				return nil, errcode.MessageNotBelongGroup
			}
			err = tx.Where("messages.to = ? AND messages.created_at > ?", request.ID, targetMsg.CreatedAt).Order("messages.created_at ASC").Limit(request.Num).Find(&msg).Error
			if err != nil {
				tx.Rollback()
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
		if request.MsgID == 0 {
			err = tx.Where("(messages.from = ? AND messages.to = ?) OR (messages.from = ? AND messages.to = ?)", uid, request.ID, request.ID, uid).Order("messages.created_at ASC").Limit(request.Num).Find(&msg).Error
			if err != nil {
				tx.Rollback()
				return nil, errcode.FindDataError.WithDetail(err.Error())
			}
		} else {
			targetMsg := entity.Message{
				Model: gorm.Model{
					ID: uint(request.MsgID),
				},
			}
			err = tx.First(&targetMsg).Error
			if err != nil {
				tx.Rollback()
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return nil, errcode.NoMessageFound
				}
				return nil, errcode.FindDataError.WithDetail(err.Error())
			}
			if (targetMsg.From != uid && targetMsg.To != request.ID && targetMsg.From != request.ID && targetMsg.To != uid) || uid == request.ID {
				tx.Rollback()
				return nil, errcode.MessageNotBelongGroup
			}
			err = tx.Where("((messages.from = ? AND messages.to = ?) OR (messages.from = ? AND messages.to = ?)) AND messages.created_at > ?", uid, request.ID, request.ID, uid, targetMsg.CreatedAt).Order("messages.created_at ASC").Limit(request.Num).Find(&msg).Error
			if err != nil {
				tx.Rollback()
				return nil, errcode.FindDataError.WithDetail(err.Error())
			}
		}
	}
	result := []response.Message{}
	for _, v := range msg {
		result = append(result, response.Message{
			ID:     int(v.ID),
			Data:   v.Data,
			Sender: v.From,
			Time:   int(v.CreatedAt.Unix()),
		})
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return nil, errcode.CommitError.WithDetail(err.Error())
	}
	return &response.Response[response.MessageDown]{
		Code:    200,
		Message: "获取成功",
		Data: &response.MessageDown{
			Msg: result,
		},
	}, nil
}
