package main

import (
	"opChat/api"
	"opChat/global"
	"opChat/http"
)

func main() {
	global.Init()

	http.Register("/user/create", false, api.User.Create)
	http.Register("/user/login", false, api.User.Login)
	http.Register("/user/setPassword", true, api.User.SetPassword)
	http.Register("/user/setName", true, api.User.SetName)

	http.Register("/group/create", true, api.Group.Create)
	http.Register("/group/agree", true, api.Group.Agree)
	http.Register("/group/delete", true, api.Group.Delete)
	http.Register("/group/setAdmin", true, api.Group.SetAdmin)
	http.Register("/group/removeAdmin", true, api.Group.RemoveAdmin)
	http.Register("/group/request", true, api.Group.Request)
	http.Register("/group/disagree", true, api.Group.Disagree)
	http.Register("/group/setName", true, api.Group.SetName)
	http.Register("/group/member", true, api.Group.Member)
	http.Register("/group/t", true, api.Group.T)

	http.Register("/organ/join", true, api.Organ.Join)
	http.Register("/organ/avatar", true, api.Organ.Avatar)
	http.Register("/organ/setAvatar", true, api.Organ.SetAvatar)
	http.Register("/organ/name", true, api.Organ.Name)
	http.Register("/organ/exit", true, api.Organ.Exit)
	http.Register("/organ/list", true, api.Organ.List)
	http.Register("/organ/avatarName", true, api.Organ.AvatarName)

	http.Register("/friend/agree", true, api.Friend.Agree)
	http.Register("/friend/disagree", true, api.Friend.Disgree)
	http.Register("/friend/request", true, api.Friend.Request)

	http.Register("/msg/send", true, api.Message.Send)
	http.Register("/msg/up", true, api.Message.Up)
	http.Register("/msg/down", true, api.Message.Down)

	e := global.StartServe()
	if e != nil {
		panic(e)
	}
}
