package v1

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zhang/LibraryMS/internal/dto"
	"github.com/zhang/LibraryMS/internal/pkg/errcode"
	"github.com/zhang/LibraryMS/internal/pkg/response"
	"github.com/zhang/LibraryMS/internal/service"
)

// UserHandler 用户处理器
type UserHandler struct {
	service service.IUserService
}

// NewUserHandler 创建用户处理器实例
func NewUserHandler(service service.IUserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

// GetByID 根据用户ID查询用户信息
// @Summary 获取用户信息
// @Description 根据用户ID获取用户详细信息
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param id path int true "用户ID"
// @Success 200 {object} response.Response
// @Router /api/v1/users/{id} [get]
func (h *UserHandler) GetByID(c *gin.Context) {
	// 1. 从请求中获取参数
	idStr := c.Param("id")
	if idStr == "" {
		response.BadRequest(c, errcode.MissingParams.Msg)
		return
	}

	// 2. 将其转换成 uint64 格式
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, errcode.InvalidParams.Msg)
		return
	}

	// 3. 调用 service 层
	user, err := h.service.FindById(id)
	if err != nil {
		if errors.Is(err, errcode.NotFound) {
			response.NotFound(c, errcode.NotFound.Msg)
			return
		}
		response.InternalError(c, errcode.InternalError.Msg)
		return
	}

	// 4. 返回成功信息（使用 DTO）
	userResp := dto.ToUserResponse(user.ID, user.Username, user.Email, uint8(user.Gender), user.Age)
	response.Success(c, userResp)
}

// Create 创建用户
// @Summary 创建用户
// @Description 创建新用户
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param user body dto.CreateUserRequest true "用户信息"
// @Success 200 {object} response.Response
// @Router /api/v1/users [post]
func (h *UserHandler) Create(c *gin.Context) {
	// 1. 从请求中获取请求参数并验证
	var req dto.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	// 2. 调用 service 层
	if err := h.service.CreateUserFromDTO(&req); err != nil {
		if errors.Is(err, errcode.UserNameDuplicated) {
			response.BadRequest(c, errcode.UserNameDuplicated.Msg)
			return
		}
		if errors.Is(err, errcode.EmailDuplicated) {
			response.BadRequest(c, errcode.EmailDuplicated.Msg)
			return
		}
		response.InternalError(c, errcode.InternalError.Msg)
		return
	}

	// 3. 返回成功信息
	response.Success(c, nil)
}

// Delete 删除用户
// @Summary 删除用户
// @Description 删除用户
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param id path int true "用户id"
// @Success 200 {boject} response.Response
// @Router /api/v1/user/{id} [delete]
func (h *UserHandler) Delete(c *gin.Context) {
	// 1.从请求中获取参数
	idStr := c.Param("id")
	if idStr == "" {
		response.BadRequest(c, errcode.MissingParams.Msg)
		return
	}

	// 2.将其转换成uint64格式
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, errcode.InvalidParams.Msg)
		return
	}

	// 3.调用 service 层
	if err := h.service.DeleteUser(id); err != nil {
		// 没有这条记录， 无法删除
		if errors.Is(err, errcode.NotFound) {
			response.NotFound(c, errcode.NotFound.Msg)
			return
		}
		// 其他错误
		response.InternalError(c, errcode.InternalError.Msg)
		return
	}

	response.Success(c, nil)
}
