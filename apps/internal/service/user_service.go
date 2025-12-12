package service

import (
	"github.com/zhang/LibraryMS/internal/model"
	"github.com/zhang/LibraryMS/internal/repository"
)

type IUserService interface {
	FindById(id uint64) (*model.User, error)
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
