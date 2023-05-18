package global

import (
	"github.com/kataras/iris/v12"
)

func Init() {
	initDatabase()
	initIris()
	initDir()
	initDefaultAvatar()
	initLogger()
	initID()
}

func StartServe() error {
	return App.Run(iris.Addr(":" + HttpPort))
}

