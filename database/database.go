package database

import (
	"context"
	"errors"
	"opChat/errcode"

	"gorm.io/gorm"
)

type database[T any] struct {
	db  *gorm.DB
}

func New[T any](db *gorm.DB, ctx context.Context) *database[T] {
	return &database[T]{
		db: db.WithContext(ctx),
	}
}

func (d *database[T]) FindByID(id uint) (*T, *errcode.Error) {
	var targetT T
	err := d.db.Find(&targetT, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errcode.NoTargetFound
		}
		return nil, errcode.FindDataError.WithDetail(err.Error())
	}
	return &targetT, nil
}

func (d *database[T]) FindByField(k string, v any) (*T, *errcode.Error) {
	var targetT T
	result := d.db.Where(k + " = ?", v).Find(&targetT)
	if result.Error != nil {
		return nil, errcode.FindDataError.WithDetail(result.Error.Error())
	}
	if result.RowsAffected == 0 {
		return nil, errcode.NoTargetFound
	}
	return &targetT, nil
}

func (d *database[T]) Add(obj *T) *errcode.Error {
	err := d.db.Create(obj).Error
	if err != nil {
		return errcode.InsertDataError.WithDetail(err.Error())
	}
	return nil
}

func (d *database[T]) Update(obj *T) *errcode.Error {
	err := d.db.Save(obj).Error
	if err != nil {
		return errcode.UpdateDataError.WithDetail(err.Error())
	}
	return nil
}

func (d *database[T]) Delete(obj *T) *errcode.Error {
	err := d.db.Delete(obj).Error
	if err != nil {
		return errcode.DeleteDataError.WithDetail(err.Error())
	}
	return nil
}
