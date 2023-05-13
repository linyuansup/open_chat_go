package http

import (
	"context"
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"opChat/database"
	"opChat/errcode"
	"opChat/global"
	"strconv"
	"strings"
	"time"

	"github.com/kataras/iris/v12"
)

func Register[req any, res any](path string, userCheck bool, action func(uid int, request *req, context context.Context) (*res, *errcode.Error)) {
	global.App.Post(path, func(c iris.Context) {
		key := ""
		userAgentFromClient := c.GetHeader("User-Agent")
		if userAgentFromClient == "" {
			errorResponse(&c, errcode.DataDecryptFail, key)
			return
		}
		ua, _ := base64.StdEncoding.DecodeString(userAgentFromClient)
		baseKey := strings.Clone(global.VersionKey)
		id := c.GetHeader("id")
		intID, err := strconv.Atoi(id)
		if err != nil {
			errorResponse(&c, errcode.TypeConventFail.WithDetail(err.Error()), key)
		}
		if userCheck {
			result, e := database.UserDatabase.FindUserByID(c.Request().Context(), intID)
			if e != nil {
				errorResponse(&c, errcode.NoUserFound, key)
				return
			}
			baseKey += result.Password + id + result.DeviceID
		}
		t := (time.Now().Unix() / 100) - 1
		path := c.Path()
		except := fmt.Sprintf("%x", md5.Sum([]byte(path)))
		for i := 0; i < 3; i++ {
			try := encrypt(ua, []byte(baseKey+fmt.Sprint(t+int64(i))))
			if string(try) == except {
				key = baseKey + fmt.Sprint(t+int64(i))
				break
			}
		}
		if key == "" {
			errorResponse(&c, errcode.DataDecryptFail, key)
			return
		}
		body, err := c.GetBody()
		if err != nil {
			errorResponse(&c, errcode.GetBodyError, key)
			return
		}
		result := encrypt(body, []byte(key))
		var unm req
		err = json.Unmarshal(result, &unm)
		if err != nil {
			errorResponse(&c, errcode.JsonFormatError, key)
			return
		}
		err = global.Validator.Struct(unm)
		if err != nil {
			errorResponse(&c, errcode.ValidatorError, key)
			return
		}
		res, e := action(intID, &unm, c.Request().Context())
		if e != nil {
			errorResponse(&c, e, key)
			return
		}
		response(&c, res, key)
	})
}

func encrypt(data, key []byte) []byte {
	d := make([]byte, len(data))
	copy(d, data)
	md5Key := md5.Sum(key)
	pos := 0
	for i := range d {
		d[i] = byte(int(d[i]) ^ int(md5Key[pos]))
		pos++
		if pos >= 16 {
			pos = 0
		}
	}
	return d
}

func errorResponse(c *iris.Context, err *errcode.Error, key string) {
	(*c).StatusCode(err.HttpCode)
	if err.HttpCode == 503 {
		return
	}
	if err.Detail != nil {
		global.Log.Error("response", err)
	}
	var result []byte
	if err.HttpCode == 502 {
		result = []byte(err.String())
	} else {
		result = encrypt([]byte(err.String()), []byte(key))
	}
	_, _ = (*c).Write(result)
}

func response(c *iris.Context, response any, key string) {
	marshal, _ := json.Marshal(response)
	marshal = encrypt(marshal, []byte(key))
	(*c).StatusCode(200)
	_, _ = (*c).Write(marshal)
}
