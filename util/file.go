package util

import (
	"crypto/md5"
	"fmt"
	"opChat/errcode"
	"opChat/global"
	"os"
)

func SaveFile(data []byte, ex string, mode string) (string, *errcode.Error) {
	fileName := fmt.Sprintf("%x", md5.Sum([]byte(data)))
	f, e := os.Create("." + global.FilePath + "/" + mode + "/" + fileName + "." + ex)
	if e != nil {
		return "", errcode.WriteDataError.WithDetail(e.Error())
	}
	defer f.Close()
	_, e = f.Write(data)
	if e != nil {
		return "", errcode.WriteDataError.WithDetail(e.Error())
	}
	return fileName, nil
}

func OpenFile(fileName string, ex string, mode string) ([]byte, *errcode.Error) {
	f, e := os.Open("." + global.FilePath + "/" + mode + "/" + fileName + "." + ex)
	if e != nil {
		return nil, errcode.OpenFileError.WithDetail(e.Error())
	}
	defer f.Close()
	info, e := f.Stat()
	if e != nil {
		return nil, errcode.OpenFileError.WithDetail(e.Error())
	}
	dst := make([]byte, info.Size())
	_, e = f.Read(dst)
	if e != nil {
		return nil, errcode.OpenFileError.WithDetail(e.Error())
	}
	return dst, nil
}
