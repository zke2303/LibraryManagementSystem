package errcode

type Error struct {
	Code int
	Msg  string
}

func (e *Error) Error() string {
	return e.Msg
}

func NewError(code int, msg string) *Error {
	return &Error{
		Code: code,
		Msg:  msg,
	}
}

// var (
//
//	// 基础错误类型
//	Success       = NewError(0, "success")
//	InvaliParams  = NewError(40001, "parameter error")
//	MissingParams = NewError(40002, "messing required parameters")
//
// )
var (
	// 基础错误
	Success       = NewError(0, "success")
	InvalidParams = NewError(40001, "参数错误")
	MissingParams = NewError(40002, "缺少必要参数")
	Unauthorized  = NewError(40003, "认证失败")
	NotFound      = NewError(40004, "资源不存在")

	// 用户模块
	UserNameEmpty  = NewError(10001, "用户名不能为空")
	PasswordEmpty  = NewError(10002, "密码不能为空")
	EmailFormatErr = NewError(10003, "邮箱格式错误")
)
