package logger

import (
	"bufio"
	"encoding/json"
	"fmt"
	"opChat/errcode"
	"os"
	"time"
)

type fileLog struct {
	writer    *bufio.Writer
	writeTime int
	filePath  string
}

func newFileLog(filePath string) (*fileLog, *errcode.Error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, errcode.OpenFileError.WithDetail(err.Error(), filePath)
	}
	filePath = dir + "/" + filePath + "/" + time.Now().Format("2006-01-02 15-04-05") + ".log"
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, errcode.OpenFileError.WithDetail(err.Error(), filePath)
	}
	return &fileLog{
		writer:    bufio.NewWriter(file),
		writeTime: 0,
		filePath:  filePath,
	}, nil
}

func formatter(level string, tag string, v ...interface{}) (string, *errcode.Error) {
	data := "{\"time\": \"" + time.Now().Format("2006-01-02 15:04:05") + "\",\"tag\":\"" + tag + "\",\"level\":\"" + level + "\",\"data\":"
	for _, ele := range v {
		marshal, err := json.Marshal(ele)
		if err != nil {
			return "", errcode.JsonFormatError.WithDetail(err.Error(), fmt.Sprint(ele))
		}
		data += string(marshal)
	}
	data += "}"
	return data, nil
}

func (f *fileLog) write(level string, tag string, v ...interface{}) *errcode.Error {
	data, err := formatter(level, tag, v)
	if err != nil {
		return err
	}
	if f.writeTime >= 500 {
		path := f.filePath + "/" + time.Now().Format("2006-01-02 15:04:05") + ".log"
		newFile, er := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
		if er != nil {
			return errcode.OpenFileError.WithDetail(er.Error(), path)
		}
		f.writer = bufio.NewWriter(newFile)
		f.writeTime = 0
	}
	_, er := f.writer.WriteString(data + "\n")
	if er != nil {
		return errcode.WriteDataError.WithDetail(er.Error())
	}
	er = f.writer.Flush()
	if err != nil {
		return errcode.WriteDataError.WithDetail(er.Error())
	}
	f.writeTime++
	return nil
}
