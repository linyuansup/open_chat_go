package global

import (
	"opChat/errcode"
	"opChat/logger"
)

var Log *logger.Logger

func initLogger() *errcode.Error {
	var err *errcode.Error
	Log, err = logger.NewLogger(logPath, log, true, true)
	return err
}
