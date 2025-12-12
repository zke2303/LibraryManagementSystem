package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/zhang/LibraryMS/internal/handler/v1"
)

func SetupBookRoutes(r *gin.RouterGroup, h *v1.BookHandler) {
	books := r.Group("/books")
	{
		books.POST("", h.Create)
	}
}
