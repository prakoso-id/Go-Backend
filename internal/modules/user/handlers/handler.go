package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go-backend/internal/modules/user/domain/service"
	"go-backend/internal/modules/user/dto"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func formatResponse(status int, message string, data interface{}, err string) Response {
	return Response{
		Status:  status,
		Message: message,
		Data:    data,
		Error:   err,
	}
}

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) Create(c *gin.Context) {
	var req dto.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, formatResponse(http.StatusBadRequest, "Invalid request", nil, err.Error()))
		return
	}

	resp, err := h.service.Create(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, formatResponse(http.StatusInternalServerError, "Failed to create user", nil, err.Error()))
		return
	}

	c.JSON(http.StatusCreated, formatResponse(http.StatusCreated, "User created successfully", resp, ""))
}

func (h *UserHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, formatResponse(http.StatusBadRequest, "Invalid ID", nil, "Invalid ID format"))
		return
	}

	resp, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, formatResponse(http.StatusNotFound, "User not found", nil, err.Error()))
		return
	}

	c.JSON(http.StatusOK, formatResponse(http.StatusOK, "User retrieved successfully", resp, ""))
}

func (h *UserHandler) List(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	resp, err := h.service.List(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, formatResponse(http.StatusInternalServerError, "Failed to retrieve users", nil, err.Error()))
		return
	}

	c.JSON(http.StatusOK, formatResponse(http.StatusOK, "Users retrieved successfully", resp, ""))
}

func (h *UserHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, formatResponse(http.StatusBadRequest, "Invalid ID", nil, "Invalid ID format"))
		return
	}

	var req dto.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, formatResponse(http.StatusBadRequest, "Invalid request", nil, err.Error()))
		return
	}

	resp, err := h.service.Update(uint(id), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, formatResponse(http.StatusInternalServerError, "Failed to update user", nil, err.Error()))
		return
	}

	c.JSON(http.StatusOK, formatResponse(http.StatusOK, "User updated successfully", resp, ""))
}

func (h *UserHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, formatResponse(http.StatusBadRequest, "Invalid ID", nil, "Invalid ID format"))
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, formatResponse(http.StatusInternalServerError, "Failed to delete user", nil, err.Error()))
		return
	}

	c.JSON(http.StatusOK, formatResponse(http.StatusOK, "User deleted successfully", nil, ""))
}

func (h *UserHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, formatResponse(http.StatusBadRequest, "Invalid request", nil, err.Error()))
		return
	}

	resp, err := h.service.Login(&req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, formatResponse(http.StatusUnauthorized, "Invalid credentials", nil, err.Error()))
		return
	}

	c.JSON(http.StatusOK, formatResponse(http.StatusOK, "User logged in successfully", resp, ""))
}

func (h *UserHandler) Logout(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, formatResponse(http.StatusUnauthorized, "User not authenticated", nil, "User not authenticated"))
		return
	}

	if err := h.service.Logout(userID.(uint)); err != nil {
		c.JSON(http.StatusInternalServerError, formatResponse(http.StatusInternalServerError, "Failed to logout", nil, err.Error()))
		return
	}

	c.JSON(http.StatusOK, formatResponse(http.StatusOK, "User logged out successfully", nil, ""))
}

func (h *UserHandler) ResetPassword(c *gin.Context) {
	var req dto.ResetPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, formatResponse(http.StatusBadRequest, "Invalid request", nil, err.Error()))
		return
	}

	resp := &dto.ResetPasswordResponse{
		Status:  "success",
		Message: "If your email exists in our system, you will receive a password reset link",
	}

	c.JSON(http.StatusOK, formatResponse(http.StatusOK, "Password reset link sent successfully", resp, ""))
}