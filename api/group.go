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

type group struct{}

var Group group

func (g *group) Create(uid int, request *request.GroupCreate) (*response.Response[response.GroupCreate], *errcode.Error) {
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
	return &response.Response[response.GroupCreate]{
		Code:    200,
		Message: "创建群聊成功",
		Data: &response.GroupCreate{
			ID: int(id),
		},
	}, nil
}

func (g *group) Delete(uid int, request *request.GroupDelete) (*response.Response[response.GroupDelete], *errcode.Error) {
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
	return &response.Response[response.GroupDelete]{
		Code:    200,
		Message: "删除群聊成功",
		Data:    &response.GroupDelete{},
	}, nil
}

func (g *group) Agree(uid int, request *request.GroupAgree) (*response.Response[response.GroupAgree], *errcode.Error) {
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
	return &response.Response[response.GroupAgree]{
		Code:    200,
		Message: "操作成功",
		Data:    &response.GroupAgree{},
	}, nil
}

func (g *group) SetAdmin(uid int, request *request.GroupSetAdmin) (*response.Response[response.GroupSetAdmin], *errcode.Error) {
	tx := global.Database.Begin()
	group := entity.Group{
		ID: uint(request.GroupID),
	}
	err := tx.First(&group).Error
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
	if uid == request.UserID {
		tx.Rollback()
		return nil, errcode.UserIsCreator
	}
	member := entity.Member{
		User:  request.UserID,
		Group: request.GroupID,
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
	if member.Admin {
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
	return &response.Response[response.GroupSetAdmin]{
		Code:    200,
		Message: "设置成功",
		Data:    &response.GroupSetAdmin{},
	}, nil
}

func (g *group) RemoveAdmin(uid int, request *request.GroupRemoveAdmin) (*response.Response[response.GroupRemoveAdmin], *errcode.Error) {
	tx := global.Database.Begin()
	group := entity.Group{
		ID: uint(request.GroupID),
	}
	err := tx.First(&group).Error
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
	member := entity.Member{
		User:  request.UserID,
		Group: request.GroupID,
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
	if !member.Admin {
		tx.Rollback()
		return nil, errcode.UserIsNotAdmin
	}
	member.Admin = false
	err = tx.Save(&member).Error
	if err != nil {
		tx.Rollback()
		return nil, errcode.UpdateDataError.WithDetail(err.Error())
	}
	err = tx.Commit().Error
	if err != nil {
		return nil, errcode.CommitError.WithDetail(err.Error())
	}
	return &response.Response[response.GroupRemoveAdmin]{
		Code:    200,
		Message: "设置成功",
		Data:    &response.GroupRemoveAdmin{},
	}, nil
}

func (g *group) Request(uid int, request *response.Request) (*response.Response[response.GroupRequest], *errcode.Error) {
	tx := global.Database.Begin()
	var memberList []entity.Member
	var groupList []entity.Group
	err := tx.Where(&entity.Group{Creator: uint(uid)}, "Creator").Find(&groupList).Error
	if err != nil {
		tx.Rollback()
		return nil, errcode.FindDataError.WithDetail(err.Error())
	}
	err = tx.Where(&entity.Member{User: uid, Admin: true}, "user", "admin").Find(&memberList).Error
	if err != nil {
		tx.Rollback()
		return nil, errcode.FindDataError.WithDetail(err.Error())
	}
	var result []response.Request
	for _, v := range memberList {
		var m []entity.Member
		err = tx.Where(&entity.Member{Grant: false, Group: v.Group}, "grant", "group").Find(&m).Error
		if err != nil {
			tx.Rollback()
			return nil, errcode.FindDataError.WithDetail(err.Error())
		}
		for _, v2 := range m {
			result = append(result, response.Request{
				ID:      v2.User,
				GroupID: v2.Group,
			})
		}
	}
	for _, v := range groupList {
		var m []entity.Member
		tx.Where(&entity.Member{Grant: false, Group: int(v.ID)}, "grant", "group").Find(&m)
		if err != nil {
			tx.Rollback()
			return nil, errcode.FindDataError.WithDetail(err.Error())
		}
		for _, v2 := range m {
			result = append(result, response.Request{
				ID:      v2.User,
				GroupID: v2.Group,
			})
		}
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return nil, errcode.CommitError.WithDetail(err.Error())
	}
	return &response.Response[response.GroupRequest]{
		Code:    200,
		Message: "获取申请列表成功",
		Data: &response.GroupRequest{
			Request: result,
		},
	}, nil
}
