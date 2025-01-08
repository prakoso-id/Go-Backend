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

func (m *Module) RegisterRoutes(router *gin.RouterGroup) {
	handler := handlers.NewHealthHandler()

	healthGroup := router.Group("/health")
	{
		healthGroup.GET("", handler.Check)
	}
}
