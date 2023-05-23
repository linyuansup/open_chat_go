package errcode

var (
	PhoneNumberAlreadyExist = NewError(20000001, "手机号已存在", 400)
	NoPhoneNumberFound      = NewError(20000002, "找不到手机号", 400)
	WrongPassword           = NewError(20000003, "密码错误", 400)
	WrongDeviceID           = NewError(20000004, "设备被更换", 400)
	NoChangePermission      = NewError(20000006, "没有修改权限", 400)
	NotInOrgan              = NewError(20000007, "不在组织内", 400)
	NoRequest               = NewError(20000008, "目标未提出申请", 400)
	NotAdmin                = NewError(20000009, "不是管理员", 400)
	NoGroupFound            = NewError(20000010, "找不到群组", 400)
	NotCreator              = NewError(20000011, "不是群主", 400)
	UserNotInGroup          = NewError(20000012, "用户不在群聊中", 400)
	UserIsCreator           = NewError(20000013, "用户是群主", 400)
	UserIsAdmin             = NewError(20000014, "用户是管理员", 400)
	AlreadyRequest          = NewError(20000015, "已存在申请", 400)
	UserIsMember            = NewError(20000016, "用户是成员", 400)
	NoUserRequestFound      = NewError(20000017, "找不到用户", 400)
	UserIsNotAdmin          = NewError(20000018, "用户不是管理员", 400)
	UserNotFriend           = NewError(20000019, "用户不是好友", 400)
)
