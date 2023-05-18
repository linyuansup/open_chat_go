package api

import (
	"context"
	"opChat/database"
	"opChat/entity"
	"opChat/errcode"
	"opChat/global"
	"opChat/request"
	"opChat/response"
)

type group struct{}

var Group group

func (g *group) Create(uid int, request *request.GroupCreateRequest, ctx context.Context) (*response.Response[response.GroupCreateResponse], *errcode.Error) {
	tx := global.Database.Begin()
	d := database.Database{
		DB: tx,
		Ctx: ctx,
	}
	group := &entity.Group{
		Name: request.Name,
		AvatarFileName: "e859977fae97b33c7e3e56d46098bd5d",
		AvatarExName: "jpg",
	}
	e := d.AddGroup(group)
	if e != nil {
		tx.Rollback()
		return nil, e
	}
	e = d.AddRelation(&entity.Relation{
		SenderID: uid,
		RecieverID: int(group.ID),
		Mode: 1,
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
		Code: 200,
		Message: "创建群聊成功",
		Data: &response.GroupCreateResponse{
			ID: int(group.ID),
		},
	}, nil
}