package router

import (
	"github.com/gin-gonic/gin"
	"go-backend/internal/modules/post"
	"go-backend/internal/modules/user"
	"gorm.io/gorm"
)

type Router struct {
	router *gin.Engine
	db     *gorm.DB
}

func NewRouter(db *gorm.DB) *Router {
	return &Router{
		router: gin.Default(),
		db:     db,
	}
}

func (r *Router) SetupRoutes() {
	// Set db in context
	r.router.Use(func(c *gin.Context) {
		c.Set("db", r.db)
		c.Next()
	})

	// API routes
	api := r.router.Group("/api")
	{
		// Register modules
		userModule := user.NewModule(r.db)
		userModule.RegisterRoutes(api)

		postModule := post.NewModule(r.db)
		postModule.RegisterRoutes(api)
	}
}

func (r *Router) Run(addr string) error {
	return r.router.Run(addr)
}
