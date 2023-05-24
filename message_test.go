package main

import (
	"opChat/api"
	"opChat/request"
	"testing"
)

func TestMessageSend(t *testing.T) {
	_, e := api.Message.Send(userID, &request.MessageSend{
		ID: userID2,
		Data: "你好",
	})
	if e != nil {
		t.Fatal("发送消息失败：" + e.Error())
	}
}
