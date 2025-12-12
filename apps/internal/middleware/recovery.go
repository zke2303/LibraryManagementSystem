package middleware

import (
	"log"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"github.com/zhang/LibraryMS/internal/pkg/response"
)

// Recovery 异常恢复中间件
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("[Recovery] panic recovered: %v\n%s", err, debug.Stack())
				response.Fail(c, http.StatusInternalServerError, http.StatusInternalServerError, "服务器内部错误")
				c.Abort()
			}
		}()
		c.Next()
	}
}
