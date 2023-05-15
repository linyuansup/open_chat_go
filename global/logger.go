package global

import (
	"opChat/errcode"
	"opChat/logger"
)

var Log *logger.Logger

func initLogger() {
	var err *errcode.Error
	Log, err = logger.NewLogger(LogPath, log, true, true)
	if err != nil {
		panic(err)
	}
}
