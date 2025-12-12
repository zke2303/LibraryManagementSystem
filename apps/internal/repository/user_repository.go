package repository

import (
	"errors"

	"github.com/go-sql-driver/mysql"
	"github.com/zhang/LibraryMS/internal/model"
	"github.com/zhang/LibraryMS/internal/pkg/errcode"
	"gorm.io/gorm"
)

type IUserRepository interface {
	FindById(id uint64) (*model.User, error)
	CreateUser(user *model.User) error
	DeleteUser(id uint64) error
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
		return nil, errcode.InternalError
	}
	return &user, nil
}

// CreateUser
// 创建一个新用户
func (repo UserRepositoryImpl) CreateUser(user *model.User) error {
	// 1.调用数据库操作
	if err := repo.db.Create(user).Error; err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return err
		}
		return errcode.InternalError
	}
	return nil
}

// DeleteUser
// delete a user
func (repo UserRepositoryImpl) DeleteUser(id uint64) error {
	// 1.调用数据库操作
	result := repo.db.Delete(&model.User{}, id)
	// 2.判断是否发生错误
	if result.Error != nil {
		return errcode.InternalError
	}

	if result.RowsAffected == 0 {
		return errcode.NotFound
	}

	return nil
}
