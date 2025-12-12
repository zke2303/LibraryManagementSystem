package service

import (
	"strings"

	"github.com/zhang/LibraryMS/internal/dto"
	"github.com/zhang/LibraryMS/internal/model"
	"github.com/zhang/LibraryMS/internal/pkg/errcode"
	"github.com/zhang/LibraryMS/internal/repository"
)

// IBookService 图书服务接口
type IBookService interface {
	CreateBook(req *dto.CreateBookRequest) error
}

// BookServiceImpl 图书服务接口实现
type BookServiceImpl struct {
	repo repository.IBookRepository
}

// NewIBookService 创建一个图书接口对象
func NewIBookService(repo repository.IBookRepository) IBookService {
	return &BookServiceImpl{
		repo: repo,
	}
}

func (service BookServiceImpl) CreateBook(req *dto.CreateBookRequest) error {
	// 1.将 dto.CreateBookRequest 对象转换成 model.Book 对象
	user := model.Book{
		Title:        req.Title,
		Author:       req.Author,
		Summary:      req.Summary,
		Price:        req.Price,
		Publisher:    req.Publisher,
		Publish_time: req.Publish_time,
		ISBN:         req.ISBN,
	}
	// 2.调用 repositoory 层
	if err := service.repo.CreateBook(&user); err != nil {
		// 判断错误类型
		if strings.Contains(err.Error(), "title") {
			return errcode.BookTitleDuplicated
		} else if strings.Contains(err.Error(), "ISBN") {
			return errcode.BookISBNDuplicated
		} else {
			return err
		}
	}

	return nil
}
