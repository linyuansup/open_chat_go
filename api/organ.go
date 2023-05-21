package api

import (
	"context"
	"opChat/database"
	"opChat/entity"
	"opChat/errcode"
	"opChat/global"
	"opChat/request"
	"opChat/response"
	"sync/atomic"

	"gorm.io/gorm"
)

type organ struct{}

var Organ organ

func (o *organ) Join(uid int, request *request.OrganJoinRequest, ctx context.Context) (*response.Response[response.OrganJoinResponse], *errcode.Error) {
	tx := global.Database.Begin()
	relationDatabase := database.New[entity.Relation](tx, ctx)
	if request.ID/100000000 >= 6 {
		groupDatabase := database.New[entity.Group](tx, ctx)
		_, e := groupDatabase.FindByID(uint(uid))
		if e != nil {
			tx.Rollback()
			if e.Code == errcode.NoTargetFound.Code {
				return nil, errcode.NoGroupFound
			}
			return nil, e
		}
	} else {
		userDatabase := database.New[entity.User](tx, ctx)
		_, e := userDatabase.FindByID(uint(request.ID))
		if e != nil {
			tx.Rollback()
			if e.Code == errcode.NoTargetFound.Code {
				return nil, errcode.NoUserFound
			}
			return nil, e
		}
	}
	result, e := relationDatabase.FindByField(&entity.Relation{
		SenderID:   uid,
		RecieverID: request.ID,
	})
	if e == nil {
		tx.Rollback()
		switch (*result)[0].Mode {
		case 0:
			{
				return nil, errcode.AlreadyRequest
			}
		case 1:
			{
				if request.ID/100000000 >= 6 {
					return nil, errcode.UserIsCreator
				}
				return nil, errcode.UserIsMember
			}
		default:
			{
				return nil, errcode.UserIsMember
			}
		}
	}
	if e.Code != errcode.NoTargetFound.Code {
		tx.Rollback()
		return nil, e
	}
	e = relationDatabase.Add(&entity.Relation{
		Model: gorm.Model{
			ID: uint(atomic.AddInt32(&global.NowRelationID, 1)),
		},
		SenderID:   uid,
		RecieverID: request.ID,
		Mode:       0,
	})
	if e != nil {
		tx.Rollback()
		return nil, e
	}
	return &response.Response[response.OrganJoinResponse]{
		Code:    200,
		Message: "申请加群成功",
		Data:    &response.OrganJoinResponse{},
	}, nil
}
