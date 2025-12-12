package errcode

var (
	// 基础错误
	Success       = NewError(0, "success")
	InvalidParams = NewError(40001, "参数错误")
	MissingParams = NewError(40002, "缺少必要参数")
	Unauthorized  = NewError(40003, "认证失败")
	NotFound      = NewError(40004, "资源不存在")
	InternalError = NewError(50000, "内部服务错误")
	// 用户模块
	UserNameEmpty  = NewError(10001, "用户名不能为空")
	PasswordEmpty  = NewError(10002, "密码不能为空")
	EmailFormatErr = NewError(10003, "邮箱格式错误")
)
