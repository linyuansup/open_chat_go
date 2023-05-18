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

	e := global.StartServe()
	if e != nil {
		panic(e)
	}
}
