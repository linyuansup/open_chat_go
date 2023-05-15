package main

import (
	"errors"
	"fmt"
	"opChat/api"
	"opChat/database"
	"opChat/entity"
	"opChat/global"
	"opChat/http"
	"os"

	"gorm.io/gorm"
)

func main() {
	initUserID()
	initDefaultAvatar()
	initDir()

	http.Register("/user/create", false, api.User.Create)
	http.Register("/user/login", false, api.User.Login)

	e := global.StartServe()
	if e != nil {
		panic(e)
	}
}

func initUserID() {
	u := entity.User{}
	err := global.Database.Last(&u).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			panic(err)
		}
		database.NowUserID = 100000000
	} else {
		database.NowUserID = int32(u.ID)
	}
	global.Log.Info("startup", fmt.Sprintf("NowUserID = %d", database.NowUserID))
}

func initDefaultAvatar() {
	wd, _ := os.Getwd()
	_, err := os.Stat(wd + "/storage/avatar/e859977fae97b33c7e3e56d46098bd5d.jpg")
	if err != nil {
		panic(err)
	}
}

func initDir() {
	dir, _ := os.Getwd()
	os.MkdirAll(dir+"/log", os.ModePerm)
}
