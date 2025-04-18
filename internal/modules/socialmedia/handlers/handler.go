package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go-backend/internal/modules/socialmedia/domain/entity"
	"go-backend/internal/modules/socialmedia/domain/service"
	"go-backend/internal/modules/socialmedia/dto"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

type SocialMediaHandler struct {
	service service.SocialMediaService
}

func NewSocialMediaHandler(service service.SocialMediaService) *SocialMediaHandler {
	return &SocialMediaHandler{service: service}
}

func formatResponse(status int, message string, data interface{}, err string) Response {
	return Response{
		Status:  status,
		Message: message,
		Data:    data,
		Error:   err,
	}
}

func (h *SocialMediaHandler) Create(c *gin.Context) {
	var req dto.CreateSocialMediaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, formatResponse(http.StatusBadRequest, "Invalid request", nil, err.Error()))
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, formatResponse(http.StatusUnauthorized, "User not authenticated", nil, "User not authenticated"))
		return
	}

	socialMedia := &entity.SocialMedia{
		Platform:  req.Platform,
		Url:       req.Url,
		ProfileID: req.ProfileID,
		UserID:    userID.(uint),
	}

	response, err := h.service.Create(socialMedia)
	if err != nil {
		c.JSON(http.StatusInternalServerError, formatResponse(http.StatusInternalServerError, "Failed to create social media", nil, err.Error()))
		return
	}

	c.JSON(http.StatusCreated, formatResponse(http.StatusCreated, "Social media created successfully", response, ""))
}

func (h *SocialMediaHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, formatResponse(http.StatusBadRequest, "Invalid ID", nil, "Invalid ID format"))
		return
	}

	response, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, formatResponse(http.StatusNotFound, "Social media not found", nil, err.Error()))
		return
	}

	c.JSON(http.StatusOK, formatResponse(http.StatusOK, "Social media retrieved successfully", response, ""))
}

func (h *SocialMediaHandler) GetAll(c *gin.Context) {
	response, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, formatResponse(http.StatusInternalServerError, "Failed to retrieve social media", nil, err.Error()))
		return
	}

	c.JSON(http.StatusOK, formatResponse(http.StatusOK, "Social media retrieved successfully", response, ""))
}

func (h *SocialMediaHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, formatResponse(http.StatusBadRequest, "Invalid ID", nil, "Invalid ID format"))
		return
	}

	var req dto.UpdateSocialMediaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, formatResponse(http.StatusBadRequest, "Invalid request", nil, err.Error()))
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, formatResponse(http.StatusUnauthorized, "User not authenticated", nil, "User not authenticated"))
		return
	}

	response, err := h.service.Update(uint(id), userID.(uint), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, formatResponse(http.StatusInternalServerError, "Failed to update social media", nil, err.Error()))
		return
	}

	c.JSON(http.StatusOK, formatResponse(http.StatusOK, "Social media updated successfully", response, ""))
}

func (h *SocialMediaHandler) Delete(c *gin.Context) {
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
		c.JSON(http.StatusInternalServerError, formatResponse(http.StatusInternalServerError, "Failed to delete social media", nil, err.Error()))
		return
	}

	c.JSON(http.StatusOK, formatResponse(http.StatusOK, "Social media deleted successfully", nil, ""))
}

func (h *SocialMediaHandler) GetByUserID(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, formatResponse(http.StatusUnauthorized, "User not authenticated", nil, "User not authenticated"))
		return
	}

	response, err := h.service.GetByUserID(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, formatResponse(http.StatusInternalServerError, "Failed to retrieve user's social media", nil, err.Error()))
		return
	}

	c.JSON(http.StatusOK, formatResponse(http.StatusOK, "User's social media retrieved successfully", response, ""))
}

func (h *SocialMediaHandler) GetByProfileID(c *gin.Context) {
	profileID, err := strconv.ParseUint(c.Param("profile_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, formatResponse(http.StatusBadRequest, "Invalid Profile ID", nil, "Invalid Profile ID format"))
		return
	}

	response, err := h.service.GetByProfileID(uint(profileID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, formatResponse(http.StatusInternalServerError, "Failed to retrieve profile's social media", nil, err.Error()))
		return
	}

	c.JSON(http.StatusOK, formatResponse(http.StatusOK, "Profile's social media retrieved successfully", response, ""))
}