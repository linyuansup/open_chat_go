package database

import (
	"context"
	"errors"
	"opChat/entity"
	"opChat/errcode"
	"opChat/global"
	"sync/atomic"

	"gorm.io/gorm"
)

type userDatabase struct{}

var UserDatabase userDatabase

func (u *userDatabase) FindByID(id uint, ctx context.Context) (*entity.User, *errcode.Error) {
	targetUser := entity.User{
		Model: gorm.Model{
			ID: id,
		},
	}
	err := global.Database.WithContext(ctx).Take(&targetUser).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errcode.NoUserFound
		}
		return nil, errcode.FindDataError.WithDetail(err.Error())
	}
	return &targetUser, nil
}

func (u *userDatabase) FindByPhoneNumber(phoneNumber string, ctx context.Context) (*entity.User, *errcode.Error) {
	targetUser := entity.User{}
	err := global.Database.WithContext(ctx).Where("phone_number = ?", phoneNumber).Take(&targetUser).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errcode.NoUserFound
		}
		return nil, errcode.FindDataError.WithDetail(err.Error())
	}
	return &targetUser, nil
}

func (u *userDatabase) Add(user *entity.User, ctx context.Context) *errcode.Error {
	user.ID = uint(atomic.AddInt32(&global.NowUserID, 1))
	err := global.Database.WithContext(ctx).Create(user).Error
	if err != nil {
		return errcode.InsertDataError.WithDetail(err.Error())
	}
	return nil
}
