package global

import (
	"github.com/go-playground/validator/v10"
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
)

var (
	Validator context.Validator
	log       *golog.Logger
	App       *iris.Application
)

func initIris() {
	App = iris.New()
	App.Use(recover.New())
	App.Use(logger.New())
	Validator = validator.New()
	log = App.Logger()
}
