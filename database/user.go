package database

import (
	"errors"
	"opChat/entity"
	"opChat/errcode"
	"opChat/global"
	"sync/atomic"

	"gorm.io/gorm"
)

func (d Database) FindUserByID(id uint) (*entity.User, *errcode.Error) {
	targetUser := entity.User{
		Model: gorm.Model{
			ID: id,
		},
	}
	err := d.DB.WithContext(d.Ctx).Take(&targetUser).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errcode.NoUserFound
		}
		return nil, errcode.FindDataError.WithDetail(err.Error())
	}
	return &targetUser, nil
}

func (d Database) FindUserByPhoneNumber(phoneNumber string) (*entity.User, *errcode.Error) {
	targetUser := entity.User{}
	err := d.DB.WithContext(d.Ctx).Where("phone_number = ?", phoneNumber).Take(&targetUser).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errcode.NoUserFound
		}
		return nil, errcode.FindDataError.WithDetail(err.Error())
	}
	return &targetUser, nil
}

func (d Database) AddUser(user *entity.User) *errcode.Error {
	user.ID = uint(atomic.AddInt32(&global.NowUserID, 1))
	err := d.DB.WithContext(d.Ctx).Create(user).Error
	if err != nil {
		return errcode.InsertDataError.WithDetail(err.Error())
	}
	return nil
}

func (d Database) UpdateUser(user *entity.User, key string, value any) *errcode.Error {
	e := d.DB.WithContext(d.Ctx).Model(user).Update(key, value).Error
	if e != nil {
		return errcode.UpdateDataError.WithDetail(e.Error())
	}
	return nil
}
