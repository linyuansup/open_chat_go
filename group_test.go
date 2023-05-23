package main

import (
	"opChat/api"
	"opChat/request"
	"testing"
)

func TestGroupCreate(t *testing.T) {
	res, e := api.Group.Create(userID, &request.GroupCreate{
		Name: "测试群组",
	})
	if e != nil {
		t.Fatal("创建群聊失败：" + e.Error())
	}
	groupID = res.Data.ID
	res, e = api.Group.Create(userID, &request.GroupCreate{
		Name: "测试群组 2",
	})
	if e != nil {
		t.Fatal("创建群聊失败：" + e.Error())
	}
	groupID2 = res.Data.ID
}

func TestGroupDelete(t *testing.T) {
	_, e := api.Group.Delete(userID2, &request.GroupDelete{
		ID: groupID2,
	})
	if e == nil {
		t.Fatal("使用非群主删除群聊成功")
	}
	_, e = api.Group.Delete(userID, &request.GroupDelete{
		ID: groupID2,
	})
	if e != nil {
		t.Fatal("删除群聊失败：" + e.Error())
	}
}

func TestGroupAgree(t *testing.T) {
	_, e := api.Group.Agree(userID, &request.GroupAgree{
		UserID:  userID2,
		GroupID: groupID,
	})
	if e != nil {
		t.Fatal("同意请求失败：" + e.Error())
	}
	_, e = api.Group.Agree(userID, &request.GroupAgree{
		UserID:  userID2,
		GroupID: groupID,
	})
	if e == nil {
		t.Fatal("同意重复请求成功")
	}
}

func TestGroupSetAdmin(t *testing.T) {
	_, e := api.Group.SetAdmin(userID, &request.GroupSetAdmin{
		UserID:  userID2,
		GroupID: groupID,
	})
	if e != nil {
		t.Fatal("设置管理员失败：" + e.Error())
	}
	_, e = api.Group.SetAdmin(userID, &request.GroupSetAdmin{
		UserID:  userID2,
		GroupID: groupID,
	})
	if e == nil {
		t.Fatal("重复设置管理员成功")
	}
}

func TestGroupRemoveAdmin(t *testing.T) {
	_, e := api.Group.RemoveAdmin(userID, &request.GroupRemoveAdmin{
		UserID:  userID2,
		GroupID: groupID,
	})
	if e != nil {
		t.Fatal("取消管理员失败：" + e.Error())
	}
	_, e = api.Group.RemoveAdmin(userID, &request.GroupRemoveAdmin{
		UserID:  userID2,
		GroupID: groupID,
	})
	if e == nil {
		t.Fatal("重复取消管理员成功")
	}
}
