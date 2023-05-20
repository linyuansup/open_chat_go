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
	e = groupDB.Delete(targetGroup)
	if e != nil {
		tx.Rollback()
		return nil, e
	}
	relation, e := relationDB.FindByField("reciever_id", targetGroup.ID)
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
