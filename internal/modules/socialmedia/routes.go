package socialmedia

import (
	"github.com/gin-gonic/gin"
	"go-backend/internal/infrastructure/middleware"
)

func (m *Module) RegisterRoutes(router *gin.RouterGroup) {
	socialMedia := router.Group("/social-media")
	{
		// Public routes
		socialMedia.GET("", m.Handler.GetAll)
		socialMedia.GET("/:id", m.Handler.GetByID)
		socialMedia.GET("/profile/:profile_id", m.Handler.GetByProfileID)

		// Protected routes
		protected := socialMedia.Use(middleware.JWTAuth(middleware.AccessToken))
		{
			protected.POST("", m.Handler.Create)
			protected.PUT("/:id", m.Handler.Update)
			protected.DELETE("/:id", m.Handler.Delete)
			protected.GET("/user", m.Handler.GetByUserID)
		}
	}
}