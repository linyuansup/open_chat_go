package global

import (
	"github.com/kataras/iris/v12"
)

func Init() {
	initDatabase()
	initIris()
	initLogger()
	initUserID()
}

func StartServe() error {
	return App.Run(iris.Addr(":" + HttpPort))
}

