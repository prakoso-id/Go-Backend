package handlers

import (
	"go-backend/internal/modules/portfolio/domain/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PortfolioHandler struct {
	service service.PortfolioService
}

func NewPortfolioHandler(service service.PortfolioService) *PortfolioHandler {
	return &PortfolioHandler{
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

// GetUserPortfolio handles retrieving a complete portfolio for a specific user
func (h *PortfolioHandler) GetUserPortfolio(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("user_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, formatResponse(http.StatusBadRequest, "Invalid User ID", nil, "Invalid User ID format"))
		return
	}

	portfolio, err := h.service.GetUserPortfolio(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, formatResponse(http.StatusInternalServerError, "Failed to retrieve user portfolio", nil, err.Error()))
		return
	}

	c.JSON(http.StatusOK, formatResponse(http.StatusOK, "User portfolio retrieved successfully", portfolio, ""))
}

// GetAllPortfolios handles retrieving summaries of all user portfolios
func (h *PortfolioHandler) GetAllPortfolios(c *gin.Context) {
	portfolios, err := h.service.GetAllPortfolios()
	if err != nil {
		c.JSON(http.StatusInternalServerError, formatResponse(http.StatusInternalServerError, "Failed to retrieve portfolios", nil, err.Error()))
		return
	}

	c.JSON(http.StatusOK, formatResponse(http.StatusOK, "Portfolios retrieved successfully", portfolios, ""))
}
