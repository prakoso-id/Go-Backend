package dto

import "time"

type CreateUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type UpdateUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"omitempty,min=6"`
}

type UserResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token  string `json:"token"`
	UserID uint   `json:"user_id"`
	Status string `json:"status"`
}

type ResetPasswordRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type ResetPasswordResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}