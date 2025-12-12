package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/zhang/LibraryMS/internal/handler/v1"
)

// SetupRouter 配置所有路由
func SetupRouter(r *gin.Engine,
	userHandler *v1.UserHandler,
	bookHandler *v1.BookHandler,
) {
	// API v1 版本
	apiV1 := r.Group("/api/v1")
	{
		SetupUserRoutes(apiV1, userHandler)
		SetupBookRoutes(apiV1, bookHandler)
	}
}
