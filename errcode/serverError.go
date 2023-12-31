package errcode

var (
	DatabaseConnectFail    = NewError(10000000, "连接数据库失败", 501)
	Base64DecodeError      = NewError(10000001, "Base64 解密失败", 501)
	OpenFileError          = NewError(10000002, "文件打开失败", 501)
	JsonFormatError        = NewError(10000003, "JSON 格式化失败", 501)
	WriteDataError         = NewError(10000004, "文件写入失败", 501)
	NeedEncrypted          = NewError(10000005, "需要加密", 502)
	NoDataFoundErrorInBody = NewError(10000006, "找不到请求体", 502)
	DataDecryptFail        = NewError(10000007, "数据解密失败", 502)
	GetBodyError           = NewError(10000008, "获取请求体失败", 501)
	ValidatorError         = NewError(10000009, "请求体校验失败", 501)
	NoPermission           = NewError(10000010, "没有权限", 501)
	TypeConventFail        = NewError(10000011, "数据转换失败", 501)
	FindDataError          = NewError(10000012, "查找数据错误", 501)
	NoUserFound            = NewError(10000013, "找不到用户", 502)
	InsertDataError        = NewError(10000014, "添加数据错误", 501)
	UpdateDataError        = NewError(10000015, "更新数据错误", 501)
	CommitError            = NewError(10000016, "提交事务错误", 501)
	DeleteDataError        = NewError(10000017, "删除错误", 501)
	NoTargetFound          = NewError(10000018, "找不到对象", 501)
	ImageDecodeError       = NewError(10000019, "图像解析失败", 501)
	UnsupportedType        = NewError(10000020, "不支持的类型", 501)
	ImageEncodeError       = NewError(10000021, "图像解析失败", 501)
)
