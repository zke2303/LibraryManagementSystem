package v1

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zhang/LibraryMS/internal/pkg/errcode"
	"github.com/zhang/LibraryMS/internal/pkg/response"
	"github.com/zhang/LibraryMS/internal/service"
)

type UserController struct {
	service service.IUserService
}

func NewUserController(service service.IUserService) UserController {
	return UserController{
		service: service,
	}
}

// FindById
// 根据用户id查询用户信息
func (ctl UserController) FindById(c *gin.Context) {
	// 1.从请求中获取参数
	idStr := c.Query("id")
	if idStr == "" {
		response.Fail(c, *errcode.MissingParams)
		return
	}
	// 2.将其转换成 uint64格式
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.Fail(c, *errcode.InvalidParams)
		return
	}

	// 3.调用 service 层
	user, err := ctl.service.FindById(id)
	if err != nil {
		if errors.Is(err, errcode.NotFound) {
			response.Fail(c, *errcode.NotFound)
			return
		}
		response.Fail(c, *errcode.InternalError)
		return
	}

	// 4.返回成功信息
	response.Success(c, user)
}
