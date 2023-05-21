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

type group struct{}

var Group group

func (g *group) Create(uid int, request *request.GroupCreateRequest, ctx context.Context) (*response.Response[response.GroupCreateResponse], *errcode.Error) {
	tx := global.Database.Begin()
	groupDB := database.New[entity.Group](tx, ctx)
	relationDB := database.New[entity.Relation](tx, ctx)
	group := &entity.Group{
		Model: gorm.Model{
			ID: uint(atomic.AddInt32(&global.NowGroupID, 1)),
		},
		Name:           request.Name,
		AvatarFileName: "e859977fae97b33c7e3e56d46098bd5d",
		AvatarExName:   "jpg",
	}
	e := groupDB.Add(group)
	if e != nil {
		tx.Rollback()
		return nil, e
	}
	e = relationDB.Add(&entity.Relation{
		Model: gorm.Model{
			ID: uint(atomic.AddInt32(&global.NowRelationID, 1)),
		},
		SenderID:   uid,
		RecieverID: int(group.ID),
		Mode:       1,
	})
	if e != nil {
		tx.Rollback()
		return nil, e
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
			ID: int(group.ID),
		},
	}, nil
}

func (g *group) Delete(uid int, request *request.GroupDeleteRequest, ctx context.Context) (*response.Response[response.GroupDeleteResponse], *errcode.Error) {
	tx := global.Database.Begin()
	groupDB := database.New[entity.Group](tx, ctx)
	relationDB := database.New[entity.Relation](tx, ctx)
	targetGroup, e := groupDB.FindByID(uint(request.ID))
	if e != nil {
		tx.Rollback()
		if e.Code == errcode.NoTargetFound.Code {
			return nil, errcode.NoGroupFound
		}
		return nil, e
	}
	_, e = relationDB.FindByField(&entity.Relation{
		SenderID:   uid,
		RecieverID: request.ID,
		Mode:       1,
	})
	if e != nil {
		tx.Rollback()
		if e.Code == errcode.NoTargetFound.Code {
			return nil, errcode.NotCreator
		}
		return nil, e
	}
	e = groupDB.Delete(targetGroup)
	if e != nil {
		tx.Rollback()
		return nil, e
	}
	relation, e := relationDB.FindByField(&entity.Relation{RecieverID: request.ID})
	if e != nil {
		tx.Rollback()
		return nil, e
	}
	for _, r := range *relation {
		e = relationDB.Delete(&r)
		if e != nil {
			tx.Rollback()
			return nil, e
		}
	}
	err := tx.Commit().Error
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

func (g *group) SetAdmin(uid int, request *request.GroupSetAdminRequest, ctx context.Context) (*response.Response[response.GroupSetAdminResponse], *errcode.Error) {
	tx := global.Database.Begin()
	relationDB := database.New[entity.Relation](tx, ctx)
	_, e := relationDB.FindByField(&entity.Relation{RecieverID: request.GroupID})
	if e != nil {
		tx.Rollback()
		if e.Code == errcode.NoTargetFound.Code {
			return nil, errcode.NoGroupFound
		}
		return nil, e
	}
	target, e := relationDB.FindByField(&entity.Relation{
		RecieverID: request.GroupID,
		SenderID:   request.UserID,
	})
	if e != nil {
		tx.Rollback()
		if e.Code == errcode.NoTargetFound.Code {
			return nil, errcode.UserNotInGroup
		}
		return nil, e
	}
	t := (*target)[0]
	switch t.Mode {
	case 0:
		{
			tx.Rollback()
			return nil, errcode.UserNotInGroup
		}
	case 1:
		{
			tx.Rollback()
			return nil, errcode.UserIsCreator
		}
	case 2:
		{
			tx.Rollback()
			return nil, errcode.UserIsAdmin
		}
	}
	t.Mode = 2
	e = relationDB.Update(&t)
	if e != nil {
		return nil, e
	}
	err := tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return nil, errcode.CommitError.WithDetail(err.Error())
	}
	return &response.Response[response.GroupSetAdminResponse]{
		Code:    200,
		Message: "设置为管理员成功",
		Data:    &response.GroupSetAdminResponse{},
	}, nil
}
