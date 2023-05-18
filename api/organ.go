package api

import (
	"context"
	"opChat/database"
	"opChat/errcode"
	"opChat/request"
	"opChat/response"
	"opChat/util"
)

type organ struct{}

var Organ organ

func (o *organ) GetAvatar(uid int, request *request.GetAvatarRequest, ctx context.Context) (*response.Response[response.GetAvatarResponse], *errcode.Error) {
	targetUser, e := database.UserDatabase.FindByID(uint(request.ID), ctx)
	if e != nil {
		if e.Code == errcode.NoUserFound.Code {
			return nil, errcode.NoTargetFound
		}
		return nil, e
	}
	file, e := util.OpenFile(targetUser.AvatarFileName, targetUser.AvatarExName, "avatar")
	if e != nil {
		return nil, e
	}
	return &response.Response[response.GetAvatarResponse]{
		Code:    200,
		Message: "获取头像成功",
		Data: &response.GetAvatarResponse{
			File: util.Base64Encode(file),
			Ex:   targetUser.AvatarExName,
		},
	}, nil
}
