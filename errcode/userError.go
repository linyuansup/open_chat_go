package errcode

var (
	NoUserFound = NewError(20000001, "找不到用户", 400)
)
