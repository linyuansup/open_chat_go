package global

import (
	"fmt"
	"github.com/kataras/iris/v12"
)

var port int

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
	return App.Run(iris.Addr(":" + fmt.Sprint(port)))
}
