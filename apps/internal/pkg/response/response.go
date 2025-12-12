package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 统一响应结构
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

// Success 成功响应
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: 0,
		Msg:  "success",
		Data: data,
	})
}

// SuccessWithMessage 成功响应（自定义消息）
func SuccessWithMessage(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: 0,
		Msg:  msg,
		Data: data,
	})
}

// BadRequest 请求参数错误 (400)
func BadRequest(c *gin.Context, msg string) {
	c.JSON(http.StatusBadRequest, Response{
		Code: http.StatusBadRequest,
		Msg:  msg,
		Data: nil,
	})
}

// Unauthorized 未授权 (401)
func Unauthorized(c *gin.Context, msg string) {
	c.JSON(http.StatusUnauthorized, Response{
		Code: http.StatusUnauthorized,
		Msg:  msg,
		Data: nil,
	})
}

// Forbidden 禁止访问 (403)
func Forbidden(c *gin.Context, msg string) {
	c.JSON(http.StatusForbidden, Response{
		Code: http.StatusForbidden,
		Msg:  msg,
		Data: nil,
	})
}

// NotFound 资源不存在 (404)
func NotFound(c *gin.Context, msg string) {
	c.JSON(http.StatusNotFound, Response{
		Code: http.StatusNotFound,
		Msg:  msg,
		Data: nil,
	})
}

// InternalError 内部服务错误 (500)
func InternalError(c *gin.Context, msg string) {
	c.JSON(http.StatusInternalServerError, Response{
		Code: http.StatusInternalServerError,
		Msg:  msg,
		Data: nil,
	})
}

// Fail 自定义错误响应
func Fail(c *gin.Context, httpCode int, bizCode int, msg string) {
	c.JSON(httpCode, Response{
		Code: bizCode,
		Msg:  msg,
		Data: nil,
	})
}
