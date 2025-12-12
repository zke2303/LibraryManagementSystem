package response

import (
	"github.com/gin-gonic/gin"
	"github.com/zhang/LibraryMS/internal/pkg/errcode"
)

type Result struct {
	Code int
	Msg  string
	Data interface{}
}

func Success(c *gin.Context, data interface{}) {
	if data == nil {
		c.JSON(200, gin.H{})
	}

	c.JSON(200, Result{
		Code: 200,
		Msg:  "success",
		Data: data,
	})
}

func Fail(c *gin.Context, err errcode.Error) {
	c.JSON(200, Result{
		Code: err.Code,
		Msg:  err.Msg,
		Data: nil,
	})
}
