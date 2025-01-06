package user

import (
	"go-backend/internal/infrastructure/middleware"
	"go-backend/internal/modules/user/domain/repository"
	"go-backend/internal/modules/user/domain/service"
	"go-backend/internal/modules/user/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.RouterGroup, db *gorm.DB) {
	repo := repository.NewUserRepository(db)
	svc := service.NewUserService(repo)
	handler := handlers.NewUserHandler(svc)

	// Auth routes
	auth := router.Group("/auth")
	{
		auth.POST("/login", handler.Login)
		auth.POST("/reset-password", handler.ResetPassword)

		// Protected auth routes - using AccessToken for logout
		protected := auth.Use(middleware.JWTAuth(middleware.AccessToken))
		protected.POST("/logout", handler.Logout)
	}

	// User routes
	users := router.Group("/users")
	{
		// Public routes

		users.GET("", handler.List)

		// Protected routes - using AccessToken for user management
		protected := users.Use(middleware.JWTAuth(middleware.AccessToken))
		{
			protected.POST("", handler.Create)
			protected.GET("/:id", handler.GetByID)
			protected.PUT("/:id", handler.Update)
			protected.DELETE("/:id", handler.Delete)
		}
	}
}
