package util

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"opChat/errcode"
	"opChat/global"
	"os"

	"github.com/nfnt/resize"
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

func Compress(buf []byte) ([]byte, *errcode.Error) {
	var width uint = 200
	var height uint = 200
	decodeBuf, layout, err := image.Decode(bytes.NewReader(buf))
	if err != nil {
		return nil, errcode.ImageDecodeError.WithDetail(err.Error())
	}
	set := resize.Resize(width, height, decodeBuf, resize.Lanczos3)
	NewBuf := bytes.Buffer{}
	switch layout {
	case "png":
		err = png.Encode(&NewBuf, set)
	case "jpeg", "jpg":
		err = jpeg.Encode(&NewBuf, set, &jpeg.Options{Quality: 80})
	default:
		return nil, errcode.UnsupportedType
	}
	if err != nil {
		return nil, errcode.ImageEncodeError.WithDetail(err.Error())
	}
	if NewBuf.Len() < len(buf) {
		buf = NewBuf.Bytes()
	}
	return buf, nil
}
