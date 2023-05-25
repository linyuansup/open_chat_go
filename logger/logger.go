package logger

import (
	"fmt"
	"opChat/errcode"
	"runtime"

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

func (l *Logger) Info(v ...interface{}) {
	_ = l.log("info", v)
}

func (l *Logger) Error(v ...interface{}) {
	_ = l.log("error", v)
}

func (l *Logger) log(level string, v ...interface{}) *errcode.Error {
	_, f, line, _ := runtime.Caller(3)
	var e *errcode.Error
	if l.toFile {
		e = l.fileLog.write(level, f+":"+fmt.Sprint(line), v)
	}
	if l.toTerminal {
		l.terminalLog.write(level, f+":"+fmt.Sprint(line), v)
	}
	return e
}
