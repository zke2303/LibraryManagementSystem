package repository

import (
	"errors"

	"github.com/zhang/LibraryMS/internal/model"
	"github.com/zhang/LibraryMS/internal/pkg/errcode"
	"gorm.io/gorm"
)

type IUserRepository interface {
	FindById(id uint64) (*model.User, error)
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

// NewIUserRepository
// 创建一个 IUserRepository 实例化对象
func NewIUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepositoryImpl{
		db: db,
	}
}

// FindById
// 根据用户id查询用户信息
func (repo UserRepositoryImpl) FindById(id uint64) (*model.User, error) {
	var user model.User
	// 1.查询数据库
	err := repo.db.First(&user, id).Error
	if err != nil {
		// 判断错误类型
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errcode.NotFound
		}
		return nil, err
	}
	return &user, nil
}
