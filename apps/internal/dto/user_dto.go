package dto

// CreateUserRequest 创建用户请求
type CreateUserRequest struct {
	Username string `json:"username" binding:"required,min=2,max=18"`
	Password string `json:"password" binding:"required,min=6,max=18"`
	Email    string `json:"email" binding:"required,email"`
	Gender   uint8  `json:"gender"`
	Age      uint8  `json:"age" binding:"required,min=1,max=150"`
}

// UpdateUserRequest 更改用户信息请求
type UpdateUserRequest struct {
	Id       uint64  `json:"id" binding:"required"`
	Username *string `json:"username" binding:"min=2,max=18"`
	Password *string `json:"password" binding:"min=6,max=18"`
	Email    *string `json:"email" binding:"email"`
	Gender   *uint8  `json:"gender"`
	Age      *uint8  `json:"age" binding:"min=1,max=150"`
}

// UserResponse 用户响应（不包含敏感信息）
type UserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Gender   uint8  `json:"gender"`
	Age      uint8  `json:"age"`
}

// ToUserResponse 将 User model 转换为 UserResponse
func ToUserResponse(id uint, username, email string, gender, age uint8) *UserResponse {
	return &UserResponse{
		ID:       id,
		Username: username,
		Email:    email,
		Gender:   gender,
		Age:      age,
	}
}
