package project

import (
	"github.com/gin-gonic/gin"
	"go-backend/internal/infrastructure/middleware"
)

func (m *Module) RegisterRoutes(router *gin.RouterGroup) {
	projects := router.Group("/projects")
	{
		// Public routes
		projects.GET("", m.Handler.GetAll)
		projects.GET("/:id", m.Handler.GetByID)

		// Protected routes
		protected := projects.Use(middleware.JWTAuth(middleware.AccessToken))
		{
			protected.POST("", m.Handler.Create)
			protected.PUT("/:id", m.Handler.Update)
			protected.DELETE("/:id", m.Handler.Delete)
		}
	}
}