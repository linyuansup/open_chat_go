package database

import (
	"opChat/entity"
	"opChat/errcode"
	"opChat/global"
	"sync/atomic"
)

func (d Database) AddGroup(group *entity.Group) *errcode.Error {
	group.ID = uint(atomic.AddInt32(&global.NowGroupID, 1))
	err := d.DB.WithContext(d.Ctx).Create(group).Error
	if err != nil {
		return errcode.InsertDataError.WithDetail(err.Error())
	}
	return nil
}