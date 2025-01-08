package router

import (
	"go-backend/internal/modules/health"
	"go-backend/internal/modules/post"
	"go-backend/internal/modules/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Router struct {
	*gin.Engine
	db *gorm.DB
}

func NewRouter(db *gorm.DB) *Router {
	engine := gin.Default()
	return &Router{
		Engine: engine,
		db:     db,
	}
}

func (r *Router) SetupRoutes() {
	// Set db in context
	r.Use(func(c *gin.Context) {
		c.Set("db", r.db)
		c.Next()
	})

	api := r.Group("/api")

	// Health check module (no auth required)
	healthModule := health.NewModule(r.db)
	healthModule.RegisterRoutes(api)

	// User module
	userModule := user.NewModule(r.db)
	userModule.RegisterRoutes(api)

	// Post module
	postModule := post.NewModule(r.db)
	postModule.RegisterRoutes(api)
}

func (r *Router) Run(addr string) error {
	return r.Engine.Run(addr)
}
