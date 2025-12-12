package service

import (
	"strings"

	"github.com/zhang/LibraryMS/internal/dto"
	"github.com/zhang/LibraryMS/internal/model"
	"github.com/zhang/LibraryMS/internal/pkg/errcode"
	"github.com/zhang/LibraryMS/internal/repository"
)

type IUserService interface {
	FindById(id uint64) (*model.User, error)
	CreateUser(user *model.User) error
	CreateUserFromDTO(req *dto.CreateUserRequest) error
	DeleteUser(id uint64) error
}

type UserServiceImpl struct {
	repo repository.IUserRepository
}

// NewIUserService 创建 IUserService 实例
func NewIUserService(repo repository.IUserRepository) IUserService {
	return &UserServiceImpl{
		repo: repo,
	}
}

// FindById 根据用户ID查询用户
func (s *UserServiceImpl) FindById(id uint64) (*model.User, error) {
	return s.repo.FindById(id)
}

// CreateUser 创建一个新用户
func (s *UserServiceImpl) CreateUser(user *model.User) error {
	err := s.repo.CreateUser(user)
	if err != nil {
		if strings.Contains(err.Error(), "username") {
			return errcode.UserNameDuplicated
		} else if strings.Contains(err.Error(), "email") {
			return errcode.EmailDuplicated
		} else {
			return err
		}

	}
	return nil
}

// CreateUserFromDTO 从 DTO 创建用户
func (s *UserServiceImpl) CreateUserFromDTO(req *dto.CreateUserRequest) error {
	user := &model.User{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
		Gender:   model.Gender(req.Gender),
		Age:      req.Age,
	}
	return s.CreateUser(user)
}

// Delete 删除一个用户
func (s *UserServiceImpl) DeleteUser(id uint64) error {
	return s.repo.DeleteUser(id)
}
