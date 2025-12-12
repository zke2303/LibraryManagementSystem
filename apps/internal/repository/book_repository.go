package repository

import (
	"errors"

	"github.com/go-sql-driver/mysql"
	"github.com/zhang/LibraryMS/internal/model"
	"github.com/zhang/LibraryMS/internal/pkg/errcode"
	"gorm.io/gorm"
)

// IBookRepository 图书数据库操作接口
type IBookRepository interface {
	CreateBook(book *model.Book) error
}

// BookRepository 图书数据库操作接口实现
type BookRepositoryImpl struct {
	db *gorm.DB
}

// NewIBookRepository 创建图书数据库操作接口实例
func NewIBookRepository(db *gorm.DB) IBookRepository {
	return &BookRepositoryImpl{
		db: db,
	}
}

// CreateBook
// 创建图书
func (repo BookRepositoryImpl) CreateBook(book *model.Book) error {
	// 1. 进行数据库操作
	var mysqlErr *mysql.MySQLError
	if err := repo.db.Create(&book).Error; err != nil {
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return err
		}
		return errcode.InternalError
	}

	return nil
}
