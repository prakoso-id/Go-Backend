package tool

import (
	"github.com/gin-gonic/gin"
	"go-backend/internal/infrastructure/middleware"
)

func (m *Module) RegisterRoutes(router *gin.RouterGroup) {
	tools := router.Group("/tools")
	{
		// Public routes
		tools.GET("", m.Handler.GetAll)
		tools.GET("/:id", m.Handler.GetByID)

		// Protected routes
		protected := tools.Use(middleware.JWTAuth(middleware.AccessToken))
		{
			protected.POST("", m.Handler.Create)
			protected.PUT("/:id", m.Handler.Update)
			protected.DELETE("/:id", m.Handler.Delete)
		}
	}
}