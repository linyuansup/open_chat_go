package errcode

var (
	PhoneNumberAlreadyExist = NewError(20000001, "手机号已存在", 400)
	NoPhoneNumberFound      = NewError(20000002, "找不到手机号", 400)
	WrongPassword           = NewError(20000003, "密码错误", 400)
	WrongDeviceID           = NewError(20000004, "设备被更换", 400)
)
