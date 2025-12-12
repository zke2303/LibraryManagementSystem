package service

import (
	"errors"

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
	UpdateUser(req *dto.UpdateUserRequest) error
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
		switch {
		case errors.Is(err, repository.ErrDuplicateKey):
			return errcode.UserNameDuplicated
		default:
			return errcode.InternalError.Wrap(err)
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
	err := s.repo.DeleteUser(id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return errcode.NotFound
		}
		if errors.Is(err, repository.ErrInternal) {
			return errcode.InternalError
		}
	}
	return nil
}

// UpdateUser 更新用户信息
func (s *UserServiceImpl) UpdateUser(req *dto.UpdateUserRequest) error {
	// 将对象转换成 map
	updates := make(map[string]interface{})
	if req.Username != nil {
		updates["username"] = *req.Username
	}
	if req.Password != nil {
		updates["password"] = *req.Password
	}
	if req.Email != nil {
		updates["email"] = *req.Email
	}
	if req.Gender != nil {
		updates["gender"] = *req.Gender
	}
	if req.Age != nil {
		updates["age"] = *req.Age
	}

	// 如何没有任何字段需要更新
	if len(updates) == 0 {
		return nil
	}

	err := s.repo.UpdateUser(req.Id, updates)
	if errors.Is(err, repository.ErrNotFound) {
		return errcode.NotFound
	}
	if errors.Is(err, repository.ErrInternal) {
		return errcode.InternalError
	}
	return nil
}
