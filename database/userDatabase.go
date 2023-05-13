package database

import (
	"context"
	"opChat/entity"
	"opChat/errcode"
)

type userDatabase struct{}

var UserDatabase userDatabase

func (u *userDatabase) FindUserByID(ctx context.Context, id int) (*entity.User, *errcode.Error) {
	return nil, nil
}
