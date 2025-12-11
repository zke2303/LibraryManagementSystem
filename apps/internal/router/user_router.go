package router

import "github.com/gin-gonic/gin"

func UserRouter(r *gin.Engine, c v1.UserController) {
	userRouter := r.Group("/user")
	{
		userRouter.GET("", c.FindById) // 根据用户id查询用户信息
	}
}
