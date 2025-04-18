package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	experienceEntity "go-backend/internal/modules/experience/domain/entity"
	postEntity "go-backend/internal/modules/post/domain/entity"
	profileEntity "go-backend/internal/modules/profile/domain/entity"
	projectEntity "go-backend/internal/modules/project/domain/entity"
	socialMediaEntity "go-backend/internal/modules/socialmedia/domain/entity"
	toolEntity "go-backend/internal/modules/tool/domain/entity"
	"gorm.io/gorm"
)

type PublicHandler struct {
	db *gorm.DB
}

func NewPublicHandler(db *gorm.DB) *PublicHandler {
	return &PublicHandler{
		db: db,
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

// GetProfiles handles retrieving all profiles
func (h *PublicHandler) GetProfiles(c *gin.Context) {
	var profiles []profileEntity.Profile
	result := h.db.Find(&profiles)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, formatResponse(http.StatusInternalServerError, "Failed to retrieve profiles", nil, result.Error.Error()))
		return
	}

	c.JSON(http.StatusOK, formatResponse(http.StatusOK, "Profiles retrieved successfully", profiles, ""))
}

// GetProfileByID handles retrieving a profile by ID
func (h *PublicHandler) GetProfileByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, formatResponse(http.StatusBadRequest, "Invalid ID", nil, "Invalid ID format"))
		return
	}

	var profile profileEntity.Profile
	result := h.db.First(&profile, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, formatResponse(http.StatusNotFound, "Profile not found", nil, result.Error.Error()))
		return
	}

	c.JSON(http.StatusOK, formatResponse(http.StatusOK, "Profile retrieved successfully", profile, ""))
}

// GetPosts handles retrieving all posts
func (h *PublicHandler) GetPosts(c *gin.Context) {
	var posts []postEntity.Post
	result := h.db.Find(&posts)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, formatResponse(http.StatusInternalServerError, "Failed to retrieve posts", nil, result.Error.Error()))
		return
	}

	c.JSON(http.StatusOK, formatResponse(http.StatusOK, "Posts retrieved successfully", posts, ""))
}

// GetPostByID handles retrieving a post by ID
func (h *PublicHandler) GetPostByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, formatResponse(http.StatusBadRequest, "Invalid ID", nil, "Invalid ID format"))
		return
	}

	var post postEntity.Post
	result := h.db.First(&post, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, formatResponse(http.StatusNotFound, "Post not found", nil, result.Error.Error()))
		return
	}

	c.JSON(http.StatusOK, formatResponse(http.StatusOK, "Post retrieved successfully", post, ""))
}

// GetProjects handles retrieving all projects
func (h *PublicHandler) GetProjects(c *gin.Context) {
	var projects []projectEntity.Project
	result := h.db.Find(&projects)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, formatResponse(http.StatusInternalServerError, "Failed to retrieve projects", nil, result.Error.Error()))
		return
	}

	c.JSON(http.StatusOK, formatResponse(http.StatusOK, "Projects retrieved successfully", projects, ""))
}

// GetProjectByID handles retrieving a project by ID
func (h *PublicHandler) GetProjectByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, formatResponse(http.StatusBadRequest, "Invalid ID", nil, "Invalid ID format"))
		return
	}

	var project projectEntity.Project
	result := h.db.First(&project, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, formatResponse(http.StatusNotFound, "Project not found", nil, result.Error.Error()))
		return
	}

	c.JSON(http.StatusOK, formatResponse(http.StatusOK, "Project retrieved successfully", project, ""))
}

// GetSocialMedia handles retrieving all social media
func (h *PublicHandler) GetSocialMedia(c *gin.Context) {
	var socialMedia []socialMediaEntity.SocialMedia
	result := h.db.Find(&socialMedia)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, formatResponse(http.StatusInternalServerError, "Failed to retrieve social media", nil, result.Error.Error()))
		return
	}

	c.JSON(http.StatusOK, formatResponse(http.StatusOK, "Social media retrieved successfully", socialMedia, ""))
}

// GetSocialMediaByID handles retrieving a social media by ID
func (h *PublicHandler) GetSocialMediaByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, formatResponse(http.StatusBadRequest, "Invalid ID", nil, "Invalid ID format"))
		return
	}

	var socialMedia socialMediaEntity.SocialMedia
	result := h.db.First(&socialMedia, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, formatResponse(http.StatusNotFound, "Social media not found", nil, result.Error.Error()))
		return
	}

	c.JSON(http.StatusOK, formatResponse(http.StatusOK, "Social media retrieved successfully", socialMedia, ""))
}

// GetTools handles retrieving all tools
func (h *PublicHandler) GetTools(c *gin.Context) {
	var tools []toolEntity.Tool
	result := h.db.Find(&tools)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, formatResponse(http.StatusInternalServerError, "Failed to retrieve tools", nil, result.Error.Error()))
		return
	}

	c.JSON(http.StatusOK, formatResponse(http.StatusOK, "Tools retrieved successfully", tools, ""))
}

// GetToolByID handles retrieving a tool by ID
func (h *PublicHandler) GetToolByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, formatResponse(http.StatusBadRequest, "Invalid ID", nil, "Invalid ID format"))
		return
	}

	var tool toolEntity.Tool
	result := h.db.First(&tool, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, formatResponse(http.StatusNotFound, "Tool not found", nil, result.Error.Error()))
		return
	}

	c.JSON(http.StatusOK, formatResponse(http.StatusOK, "Tool retrieved successfully", tool, ""))
}

// GetExperiences handles retrieving all experiences
func (h *PublicHandler) GetExperiences(c *gin.Context) {
	var experiences []experienceEntity.Experience
	result := h.db.Find(&experiences)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, formatResponse(http.StatusInternalServerError, "Failed to retrieve experiences", nil, result.Error.Error()))
		return
	}

	c.JSON(http.StatusOK, formatResponse(http.StatusOK, "Experiences retrieved successfully", experiences, ""))
}

// GetExperienceByID handles retrieving an experience by ID
func (h *PublicHandler) GetExperienceByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, formatResponse(http.StatusBadRequest, "Invalid ID", nil, "Invalid ID format"))
		return
	}

	var experience experienceEntity.Experience
	result := h.db.First(&experience, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, formatResponse(http.StatusNotFound, "Experience not found", nil, result.Error.Error()))
		return
	}

	c.JSON(http.StatusOK, formatResponse(http.StatusOK, "Experience retrieved successfully", experience, ""))
}

// GetPortfolio handles retrieving a complete portfolio for a user
func (h *PublicHandler) GetPortfolio(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("user_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, formatResponse(http.StatusBadRequest, "Invalid User ID", nil, "Invalid User ID format"))
		return
	}

	// Get user profile
	var profile profileEntity.Profile
	result := h.db.Where("user_id = ?", userID).First(&profile)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, formatResponse(http.StatusNotFound, "Profile not found", nil, result.Error.Error()))
		return
	}

	// Get user posts
	var posts []postEntity.Post
	h.db.Where("user_id = ?", userID).Find(&posts)

	// Get user projects
	var projects []projectEntity.Project
	h.db.Where("user_id = ?", userID).Find(&projects)

	// Get user social media
	var socialMedia []socialMediaEntity.SocialMedia
	h.db.Where("user_id = ?", userID).Find(&socialMedia)

	// Get user tools
	var tools []toolEntity.Tool
	h.db.Where("user_id = ?", userID).Find(&tools)

	// Get user experiences
	var experiences []experienceEntity.Experience
	h.db.Where("user_id = ?", userID).Find(&experiences)

	// Create portfolio response
	portfolio := gin.H{
		"profile":      profile,
		"posts":        posts,
		"projects":     projects,
		"social_media": socialMedia,
		"tools":        tools,
		"experiences":  experiences,
	}

	c.JSON(http.StatusOK, formatResponse(http.StatusOK, "Portfolio retrieved successfully", portfolio, ""))
}
