package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go-backend/internal/modules/post/domain/service"
	"go-backend/internal/modules/post/dto"
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

type PostHandler struct {
	service service.PostService
}

func NewPostHandler(service service.PostService) *PostHandler {
	return &PostHandler{service: service}
}

func (h *PostHandler) Create(c *gin.Context) {
	var req dto.CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, formatResponse(http.StatusBadRequest, "Invalid request", nil, err.Error()))
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, formatResponse(http.StatusUnauthorized, "User not authenticated", nil, "User not authenticated"))
		return
	}

	resp, err := h.service.Create(userID.(uint), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, formatResponse(http.StatusInternalServerError, "Failed to create post", nil, err.Error()))
		return
	}

	c.JSON(http.StatusCreated, formatResponse(http.StatusCreated, "Post created successfully", resp, ""))
}

func (h *PostHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, formatResponse(http.StatusBadRequest, "Invalid ID", nil, "Invalid ID format"))
		return
	}

	resp, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, formatResponse(http.StatusNotFound, "Post not found", nil, err.Error()))
		return
	}

	c.JSON(http.StatusOK, formatResponse(http.StatusOK, "Post retrieved successfully", resp, ""))
}

func (h *PostHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, formatResponse(http.StatusBadRequest, "Invalid ID", nil, "Invalid ID format"))
		return
	}

	var req dto.UpdatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, formatResponse(http.StatusBadRequest, "Invalid request", nil, err.Error()))
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, formatResponse(http.StatusUnauthorized, "User not authenticated", nil, "User not authenticated"))
		return
	}

	resp, err := h.service.Update(uint(id), userID.(uint), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, formatResponse(http.StatusInternalServerError, "Failed to update post", nil, err.Error()))
		return
	}

	c.JSON(http.StatusOK, formatResponse(http.StatusOK, "Post updated successfully", resp, ""))
}

func (h *PostHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, formatResponse(http.StatusBadRequest, "Invalid ID", nil, "Invalid ID format"))
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, formatResponse(http.StatusUnauthorized, "User not authenticated", nil, "User not authenticated"))
		return
	}

	if err := h.service.Delete(uint(id), userID.(uint)); err != nil {
		c.JSON(http.StatusInternalServerError, formatResponse(http.StatusInternalServerError, "Failed to delete post", nil, err.Error()))
		return
	}

	c.JSON(http.StatusOK, formatResponse(http.StatusOK, "Post deleted successfully", nil, ""))
}

func (h *PostHandler) List(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	posts, err := h.service.List(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, formatResponse(http.StatusInternalServerError, "Failed to retrieve posts", nil, err.Error()))
		return
	}

	c.JSON(http.StatusOK, formatResponse(http.StatusOK, "Posts retrieved successfully", posts, ""))
}

func (h *PostHandler) ListByUserID(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("user_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, formatResponse(http.StatusBadRequest, "Invalid user ID", nil, "Invalid ID format"))
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	posts, err := h.service.ListByUserID(uint(userID), page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, formatResponse(http.StatusInternalServerError, "Failed to retrieve user's posts", nil, err.Error()))
		return
	}

	c.JSON(http.StatusOK, formatResponse(http.StatusOK, "User's posts retrieved successfully", posts, ""))
}