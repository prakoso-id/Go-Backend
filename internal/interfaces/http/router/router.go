package router

import (
	"go-backend/internal/modules/experience"
	"go-backend/internal/modules/health"
	"go-backend/internal/modules/post"
	"go-backend/internal/modules/profile"
	"go-backend/internal/modules/project"
	"go-backend/internal/modules/public"
	"go-backend/internal/modules/socialmedia"
	"go-backend/internal/modules/tool"
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
	
	// Enable CORS
	engine.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		
		c.Next()
	})

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

	// Project module
	projectModule := project.NewModule(r.db)
	projectModule.RegisterRoutes(api)

	// Tool module
	toolModule := tool.NewModule(r.db)
	toolModule.RegisterRoutes(api)

	// Profile module
	profileModule := profile.NewModule(r.db)
	profileModule.RegisterRoutes(api)

	// Social Media module
	socialMediaModule := socialmedia.NewModule(r.db)
	socialMediaModule.RegisterRoutes(api)

	// Experience module
	experienceModule := experience.NewModule(r.db)
	experienceModule.RegisterRoutes(api)

	// Public API module
	publicModule := public.NewModule(r.db)
	publicModule.RegisterRoutes(api)
}

func (r *Router) Run(addr string) error {
	return r.Engine.Run(addr)
}
