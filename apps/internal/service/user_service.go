package service

import (
	"errors"

	"github.com/zhang/LibraryMS/internal/model"
	"github.com/zhang/LibraryMS/internal/pkg/errcode"
	"github.com/zhang/LibraryMS/internal/repository"
)

type IUserService interface {
	FindById(id uint64) (*model.User, error)
	CreateUser(user *model.User) error
}

type UserServiceImpl struct {
	repo repository.IUserRepository
}

// NewIUserService
// 创建 IUserSservice 示例
func NewIUserService(repo repository.IUserRepository) IUserService {
	return &UserServiceImpl{
		repo: repo,
	}
}

// FindById
// 根据用户id查询用户
func (service *UserServiceImpl) FindById(id uint64) (*model.User, error) {
	return service.repo.FindById(id)
}

// CreateUser
// 创建一个新用户
func (service *UserServiceImpl) CreateUser(user *model.User) error {
	err := service.repo.CreateUser(user)
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrDuplicateKey):
			return errcode.UserNameDuplicated
		default:
			return errcode.InternalError.Wrap(err)
		}
	}
	return nil
}
