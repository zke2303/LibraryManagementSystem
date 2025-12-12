package v1

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zhang/LibraryMS/internal/model"
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

func (ctl UserController) CreateUser(c *gin.Context) {
	// 1.从请求中获取请求参数
	var user model.User
	if err := c.BindJSON(&user); err != nil {
		// TODO: 错误分析
		// 判断请求参数是否正确
		fmt.Println(err)
		response.Fail(c, *errcode.InvalidParams)
		return
	}

	// 2.调用 service 层
	if err := ctl.service.CreateUser(&user); err != nil {
		// TODO: 错误分析
		response.Fail(c, *errcode.UserNameDuplicated)
		return
	}

	// 3.返回成功信息
	response.Success(c, nil)
}
