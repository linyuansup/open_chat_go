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
	"opChat/response"
	"strconv"
	"strings"
	"time"

	"github.com/kataras/iris/v12"
)

func Register[req any, res any](path string, userCheck bool, action func(uid int, request *req, ctx context.Context) (*response.Response[res], *errcode.Error)) {
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
			result, e := database.Database{
				DB:  global.Database,
				Ctx: c.Request().Context(),
			}.FindUserByID(uint(intID))
			if e != nil {
				errorResponse(&c, errcode.NoUserFound, key)
				return
			}
			baseKey += result.Password + id + result.DeviceID
		}
		t := (time.Now().Unix() / 100) - 1
		path := c.Path()
		except := md5.Sum([]byte(path))
		for i := 0; i < 3; i++ {
			try := encrypt(ua, []byte(baseKey+fmt.Sprint(t+int64(i))))
			if compare(try, except) {
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
		err = global.Validator.Struct(&unm)
		if err != nil {
			errorResponse(&c, errcode.ValidatorError, key)
			return
		}
		res, e := action(intID, &unm, c.Request().Context())
		if e != nil {
			errorResponse(&c, e, key)
			return
		}
		successResponse(&c, res, key)
	})
}

func encrypt(data, key []byte) []byte {
	if len(key) == 0 || len(data) == 0 {
		return data
	}
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
	var result []byte
	if err.HttpCode == 400 {
		result = []byte(fmt.Sprintf("{\"code\":%d,\"message\":\"%s\",\"data\":{}}", err.Code, err.Message))
	} else {
		if err.HttpCode == 501 {
			result = []byte(fmt.Sprintf("{\"code\":%d,\"message\":\"\",\"data\":{}}", err.Code))
		}
	}
	global.Log.Info("err_response", err)
	_, _ = (*c).Write(encrypt(result, []byte(key)))
}

func successResponse(c *iris.Context, response any, key string) {
	marshal, err := json.Marshal(response)
	if err != nil {
		errorResponse(c, errcode.JsonFormatError.WithDetail(err.Error()), key)
		return
	}
	(*c).StatusCode(200)
	global.Log.Info("success_response", marshal)
	_, _ = (*c).Write(encrypt(marshal, []byte(key)))
}

func compare(a []byte, b [16]byte) bool {
	if len(a) != 16 {
		return false
	}
	for i, v := range b {
		if a[i] != v {
			return false
		}
	}
	return true
}
