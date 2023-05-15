package errcode

var (
	PhoneNumberAlreadyExist = NewError(20000001, "手机号已存在", 400)
)
