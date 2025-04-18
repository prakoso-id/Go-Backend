package handlers

import (
	"net/http"
	"strconv"

	imageEntity "go-backend/internal/modules/images/domain/entity"
	"go-backend/internal/modules/project/domain/entity"
	"go-backend/internal/modules/project/domain/service"
	"go-backend/internal/modules/project/dto"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

type ProjectHandler struct {
	service service.ProjectService
	db      *gorm.DB
}

func NewProjectHandler(service service.ProjectService, db *gorm.DB) *ProjectHandler {
	return &ProjectHandler{
		service: service,
		db:      db,
	}
}

func formatResponse(status int, message string, data interface{}, err string) Response {
	return Response{
		Status:  status,
		Message: message,
		Data:    data,
		Error:   err,
	}
}

func (h *ProjectHandler) Create(c *gin.Context) {
	var req dto.CreateProjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, formatResponse(http.StatusBadRequest, "Invalid request", nil, err.Error()))
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, formatResponse(http.StatusUnauthorized, "User not authenticated", nil, "User not authenticated"))
		return
	}

	project := &entity.Project{
		Name:        req.Name,
		Description: req.Description,
		Url:         req.Url,
		UserID:      userID.(uint),
	}

	// Create images if provided
	if len(req.ImageURLs) > 0 {
		images := make([]imageEntity.Images, len(req.ImageURLs))
		for i, url := range req.ImageURLs {
			projectID := project.ID
			images[i] = imageEntity.Images{
				URL:       url,
				ProjectID: &projectID,
				UserID:    userID.(uint),
			}
		}
		project.Images = images
	}

	resp, err := h.service.Create(project)
	if err != nil {
		c.JSON(http.StatusInternalServerError, formatResponse(http.StatusInternalServerError, "Failed to create project", nil, err.Error()))
		return
	}

	c.JSON(http.StatusCreated, formatResponse(http.StatusCreated, "Project created successfully", resp, ""))
}

func (h *ProjectHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, formatResponse(http.StatusBadRequest, "Invalid ID", nil, "Invalid ID format"))
		return
	}

	resp, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, formatResponse(http.StatusNotFound, "Project not found", nil, err.Error()))
		return
	}

	c.JSON(http.StatusOK, formatResponse(http.StatusOK, "Project retrieved successfully", resp, ""))
}

func (h *ProjectHandler) GetAll(c *gin.Context) {
	response, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, formatResponse(http.StatusInternalServerError, "Failed to retrieve projects", nil, err.Error()))
		return
	}

	c.JSON(http.StatusOK, formatResponse(http.StatusOK, "Projects retrieved successfully", response, ""))
}

func (h *ProjectHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, formatResponse(http.StatusBadRequest, "Invalid ID", nil, "Invalid ID format"))
		return
	}

	var req dto.UpdateProjectRequest
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
		c.JSON(http.StatusInternalServerError, formatResponse(http.StatusInternalServerError, "Failed to update project", nil, err.Error()))
		return
	}

	c.JSON(http.StatusOK, formatResponse(http.StatusOK, "Project updated successfully", resp, ""))
}

func (h *ProjectHandler) Delete(c *gin.Context) {
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
		c.JSON(http.StatusInternalServerError, formatResponse(http.StatusInternalServerError, "Failed to delete project", nil, err.Error()))
		return
	}

	c.JSON(http.StatusOK, formatResponse(http.StatusOK, "Project deleted successfully", nil, ""))
}

func (h *ProjectHandler) GetByUserID(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, formatResponse(http.StatusUnauthorized, "User not authenticated", nil, "User not authenticated"))
		return
	}

	response, err := h.service.GetByUserID(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, formatResponse(http.StatusInternalServerError, "Failed to retrieve user's projects", nil, err.Error()))
		return
	}

	c.JSON(http.StatusOK, formatResponse(http.StatusOK, "User's projects retrieved successfully", response, ""))
}
