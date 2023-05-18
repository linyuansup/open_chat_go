package database

import (
	"opChat/entity"
	"opChat/errcode"
	"opChat/global"
	"sync/atomic"
)

func (d Database) AddRelation(relation *entity.Relation) *errcode.Error {
	relation.ID = uint(atomic.AddInt32(&global.NowRelationID, 1))
	err := d.DB.WithContext(d.Ctx).Create(relation).Error
	if err != nil {
		return errcode.InsertDataError.WithDetail(err.Error())
	}
	return nil
}