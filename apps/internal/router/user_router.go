package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/zhang/LibraryMS/internal/api/v1"
)

func UserRouter(r *gin.Engine, c v1.UserController) {
	userRouter := r.Group("/api/user")
	{
		userRouter.GET("", c.FindById) // 根据用户id查询用户信息
		userRouter.POST("", c.CreateUser)
	}
}
