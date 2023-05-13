package logger

import (
	"opChat/errcode"

	"github.com/kataras/golog"
)

type Logger struct {
	*fileLog
	*terminalLog
	toFile     bool
	toTerminal bool
}

func NewLogger(filePath string, irisLogger *golog.Logger, toFile bool, toTerminal bool) (*Logger, *errcode.Error) {
	fl, err := newFileLog(filePath)
	if err != nil {
		return nil, err
	}
	tl := newTerminalLog(irisLogger)
	return &Logger{
		fileLog:     fl,
		terminalLog: tl,
		toFile:      toFile,
		toTerminal:  toTerminal,
	}, nil
}

func (l *Logger) Info(tag string, v ...interface{}) {
	_ = l.log("info", tag, v)
}

func (l *Logger) Error(tag string, v ...interface{}) {
	_ = l.log("error", tag, v)
}

func (l *Logger) log(level string, tag string, v ...interface{}) *errcode.Error {
	var e *errcode.Error
	if l.toFile {
		e = l.fileLog.write(level, tag, v)
	}
	if l.toTerminal {
		l.terminalLog.write(level, tag, v)
	}
	return e
}
