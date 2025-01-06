package post

import (
	"github.com/gin-gonic/gin"
	"go-backend/internal/infrastructure/middleware"
	"go-backend/internal/modules/post/domain/repository"
	"go-backend/internal/modules/post/domain/service"
	"go-backend/internal/modules/post/handlers"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.RouterGroup, db *gorm.DB) {
	repo := repository.NewPostRepository(db)
	svc := service.NewPostService(repo)
	handler := handlers.NewPostHandler(svc)

	posts := router.Group("/posts")
	{
		// Public routes
		posts.GET("", handler.List)
		posts.GET("/:id", handler.GetByID)
		posts.GET("/user/:user_id", handler.ListByUserID)

		// Protected routes
		protected := posts.Use(middleware.JWTAuth(middleware.AccessToken))
		{
			protected.POST("", handler.Create)
			protected.PUT("/:id", handler.Update)
			protected.DELETE("/:id", handler.Delete)
		}
	}
}