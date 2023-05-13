package util

import (
	"encoding/base64"
	"opChat/errcode"
)

func Base64Decode(src []byte) ([]byte, *errcode.Error) {
	res := make([]byte, base64.StdEncoding.DecodedLen(len(src)))
	_, err := base64.StdEncoding.Decode(res, src)
	if err != nil {
		return nil, errcode.Base64DecodeError.WithDetail(err.Error())
	}
	return res, nil
}

func Base64Encode(src []byte) string {
	res := make([]byte, base64.StdEncoding.EncodedLen(len(src)))
	base64.StdEncoding.Encode(res, src)
	return string(res)
}
