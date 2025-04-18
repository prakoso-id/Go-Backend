package handlers

import (
	"go-backend/internal/modules/experience/domain/service"
	"go-backend/internal/modules/experience/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ExperienceHandler struct {
	service service.ExperienceService
}

func NewExperienceHandler(service service.ExperienceService) *ExperienceHandler {
	return &ExperienceHandler{
		service: service,
	}
}

// formatResponse formats the response to a standard format
func formatResponse(statusCode int, message string, data interface{}, error string) gin.H {
	return gin.H{
		"status":  statusCode,
		"message": message,
		"data":    data,
		"error":   error,
	}
}

// Create handles the creation of a new experience
func (h *ExperienceHandler) Create(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, formatResponse(http.StatusUnauthorized, "User not authenticated", nil, "User not authenticated"))
		return
	}

	var request dto.CreateExperienceRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, formatResponse(http.StatusBadRequest, "Invalid request body", nil, err.Error()))
		return
	}

	response, err := h.service.Create(&request, userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, formatResponse(http.StatusInternalServerError, "Failed to create experience", nil, err.Error()))
		return
	}

	c.JSON(http.StatusCreated, formatResponse(http.StatusCreated, "Experience created successfully", response, ""))
}

// GetAll handles retrieving all experiences
func (h *ExperienceHandler) GetAll(c *gin.Context) {
	response, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, formatResponse(http.StatusInternalServerError, "Failed to retrieve experiences", nil, err.Error()))
		return
	}

	c.JSON(http.StatusOK, formatResponse(http.StatusOK, "Experiences retrieved successfully", response, ""))
}

// GetByID handles retrieving an experience by ID
func (h *ExperienceHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, formatResponse(http.StatusBadRequest, "Invalid ID", nil, "Invalid ID format"))
		return
	}

	response, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, formatResponse(http.StatusNotFound, "Experience not found", nil, err.Error()))
		return
	}

	c.JSON(http.StatusOK, formatResponse(http.StatusOK, "Experience retrieved successfully", response, ""))
}

// GetByUserID handles retrieving experiences by user ID
func (h *ExperienceHandler) GetByUserID(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, formatResponse(http.StatusUnauthorized, "User not authenticated", nil, "User not authenticated"))
		return
	}

	response, err := h.service.GetByUserID(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, formatResponse(http.StatusInternalServerError, "Failed to retrieve user's experiences", nil, err.Error()))
		return
	}

	c.JSON(http.StatusOK, formatResponse(http.StatusOK, "User's experiences retrieved successfully", response, ""))
}



// Update handles updating an experience
func (h *ExperienceHandler) Update(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, formatResponse(http.StatusUnauthorized, "User not authenticated", nil, "User not authenticated"))
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, formatResponse(http.StatusBadRequest, "Invalid ID", nil, "Invalid ID format"))
		return
	}

	// Get the experience to verify ownership
	experience, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, formatResponse(http.StatusNotFound, "Experience not found", nil, err.Error()))
		return
	}

	// Verify that the user owns this experience
	if experience.UserID != userID.(uint) {
		c.JSON(http.StatusForbidden, formatResponse(http.StatusForbidden, "Not authorized to update this experience", nil, "You can only update your own experiences"))
		return
	}

	var request dto.UpdateExperienceRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, formatResponse(http.StatusBadRequest, "Invalid request body", nil, err.Error()))
		return
	}

	response, err := h.service.Update(uint(id), &request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, formatResponse(http.StatusInternalServerError, "Failed to update experience", nil, err.Error()))
		return
	}

	c.JSON(http.StatusOK, formatResponse(http.StatusOK, "Experience updated successfully", response, ""))
}

// Delete handles deleting an experience
func (h *ExperienceHandler) Delete(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, formatResponse(http.StatusUnauthorized, "User not authenticated", nil, "User not authenticated"))
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, formatResponse(http.StatusBadRequest, "Invalid ID", nil, "Invalid ID format"))
		return
	}

	// Get the experience to verify ownership
	experience, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, formatResponse(http.StatusNotFound, "Experience not found", nil, err.Error()))
		return
	}

	// Verify that the user owns this experience
	if experience.UserID != userID.(uint) {
		c.JSON(http.StatusForbidden, formatResponse(http.StatusForbidden, "Not authorized to delete this experience", nil, "You can only delete your own experiences"))
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, formatResponse(http.StatusInternalServerError, "Failed to delete experience", nil, err.Error()))
		return
	}

	c.JSON(http.StatusOK, formatResponse(http.StatusOK, "Experience deleted successfully", nil, ""))
}
