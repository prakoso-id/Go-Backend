package profile

import (
	"github.com/gin-gonic/gin"
	"go-backend/internal/infrastructure/middleware"
)

func (m *Module) RegisterRoutes(router *gin.RouterGroup) {
	profiles := router.Group("/profiles")
	{
		// Public routes
		profiles.GET("", m.Handler.GetAll)
		profiles.GET("/:id", m.Handler.GetByID)

		// Protected routes
		protected := profiles.Use(middleware.JWTAuth(middleware.AccessToken))
		{
			protected.POST("", m.Handler.Create)
			protected.PUT("/:id", m.Handler.Update)
			protected.DELETE("/:id", m.Handler.Delete)
		}
	}
}