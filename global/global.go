package global

import (
	"github.com/kataras/iris/v12"
)

func init() {
	err := initDatabase()
	if err != nil {
		panic(err)
	}
	err = initIris()
	if err != nil {
		panic(err)
	}
	err = initLogger()
	if err != nil {
		panic(err)
	}
}

func StartServe() error {
	return App.Run(iris.Addr(":" + HttpPort))
}
