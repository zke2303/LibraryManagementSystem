package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/zhang/LibraryMS/internal/handler/v1"
)

// SetupUserRoutes 配置用户相关路由
func SetupUserRoutes(rg *gin.RouterGroup, h *v1.UserHandler) {
	users := rg.Group("/users")
	{
		users.GET("/:id", h.GetByID)   // 根据用户ID查询用户信息
		users.POST("", h.Create)       // 创建用户
		users.DELETE("/:id", h.Delete) // 删除用户
		users.PUT("", h.Update)
	}
}
