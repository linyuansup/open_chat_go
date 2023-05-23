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
	http.Register("/group/setAdmin", true, api.Group.SetAdmin)
	http.Register("/group/request", true, api.Group.Request)
	http.Register("/group/disagree", true, api.Group.Disagree)
	http.Register("/group/setName", true, api.Group.SetName)

	http.Register("/organ/join", true, api.Organ.Join)
	http.Register("/organ/avatar", true, api.Organ.Avatar)
	http.Register("/organ/setAvatar", true, api.Organ.SetAvatar)
	http.Register("/organ/name", true, api.Organ.Name)
	http.Register("/organ/exut", true, api.Organ.Exit)

	http.Register("/friend/agree", true, api.Friend.Agree)

	e := global.StartServe()
	if e != nil {
		panic(e)
	}
}
