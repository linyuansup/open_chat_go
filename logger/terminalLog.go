package logger

import (
	"github.com/kataras/golog"
)

type terminalLog struct {
	logger *golog.Logger
}

func newTerminalLog(logger *golog.Logger) *terminalLog {
	return &terminalLog{logger: logger}
}

func (t *terminalLog) write(level string, tag string, v ...interface{}) {
	if level == "info" {
		s, _ := formatter(tag, v...)
		t.logger.Info(s)
	} else {
		s, _ := formatter(tag, v...)
		t.logger.Error(s)
	}
}
