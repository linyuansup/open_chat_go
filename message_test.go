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

func TestMessageUp(t *testing.T) {
	_, e := api.Message.Up(userID, &request.MessageUp{
		ID: userID2,
		MsgID: 0,
		Num: 20,
	})
	if e != nil {
		t.Fatal("获取消息失败：" + e.Error())
	}
}

func TestMessageDown(t *testing.T) {
	_, e := api.Message.Down(userID, &request.MessageDown{
		ID: userID2,
		MsgID: 0,
		Num: 20,
	})
	if e != nil {
		t.Fatal("获取消息失败：" + e.Error())
	}
}
