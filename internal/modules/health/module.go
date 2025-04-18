package health

import (
	"go-backend/internal/modules/health/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Module struct {
	db *gorm.DB
}

func NewModule(db *gorm.DB) *Module {
	return &Module{db: db}
}

// RegisterRoutes registers the health check routes
func (m *Module) RegisterRoutes(router *gin.RouterGroup) {
	handler := handlers.NewHealthHandler()

	// Register directly to the router, not in a group
	router.GET("/health", handler.Check)
}
