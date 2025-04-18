package experience

import (
	"github.com/gin-gonic/gin"
	"go-backend/internal/infrastructure/middleware"
)

func (m *Module) RegisterRoutes(router *gin.RouterGroup) {
	experiences := router.Group("/experiences")
	{
		// Public routes
		experiences.GET("", m.Handler.GetAll)
		experiences.GET("/:id", m.Handler.GetByID)

		// Protected routes
		protected := experiences.Use(middleware.JWTAuth(middleware.AccessToken))
		{
			protected.POST("", m.Handler.Create)
			protected.PUT("/:id", m.Handler.Update)
			protected.DELETE("/:id", m.Handler.Delete)
			protected.GET("/user", m.Handler.GetByUserID)
		}
	}
}
