package v1

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/zhang/LibraryMS/internal/dto"
	"github.com/zhang/LibraryMS/internal/pkg/errcode"
	"github.com/zhang/LibraryMS/internal/pkg/response"
	"github.com/zhang/LibraryMS/internal/service"
)

// BookHandler 图书处理器
type BookHandler struct {
	service service.IBookService
}

// NewBookHandler 创建图书处理器实例
func NewBookHandler(service service.IBookService) *BookHandler {
	return &BookHandler{
		service: service,
	}
}

func (h *BookHandler) Create(c *gin.Context) {
	// 1.从请求中获取参数，并进行校验
	var req dto.CreateBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	// 2.调用 service 层
	if err := h.service.CreateBook(&req); err != nil {
		// TODO: 进行错误判断
		if errors.Is(err, errcode.BookTitleDuplicated) {
			response.BadRequest(c, errcode.BookTitleDuplicated.Msg)
			return
		}
		if errors.Is(err, errcode.BookISBNDuplicated) {
			response.BadRequest(c, errcode.BookISBNDuplicated.Msg)
			return
		}
		response.InternalError(c, errcode.InternalError.Msg)
		return
	}

	// 3.返回成功信息
	response.Success(c, nil)
}
