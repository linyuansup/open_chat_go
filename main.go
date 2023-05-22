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

	http.Register("/group/create", true, api.Group.Create)

	http.Register("/organ/join", true, api.Oran.Join)

	e := global.StartServe()
	if e != nil {
		panic(e)
	}
}
