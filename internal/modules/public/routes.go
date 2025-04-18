package public

import (
	"github.com/gin-gonic/gin"
)

func (m *Module) RegisterRoutes(router *gin.RouterGroup) {
	// Create a public API group
	public := router.Group("/public")
	{
		// Portfolio endpoint - gets everything for a user
		public.GET("/portfolio/:user_id", m.Handler.GetPortfolio)

		// Individual resource endpoints
		profiles := public.Group("/profiles")
		{
			profiles.GET("", m.Handler.GetProfiles)
			profiles.GET("/:id", m.Handler.GetProfileByID)
		}

		posts := public.Group("/posts")
		{
			posts.GET("", m.Handler.GetPosts)
			posts.GET("/:id", m.Handler.GetPostByID)
		}

		projects := public.Group("/projects")
		{
			projects.GET("", m.Handler.GetProjects)
			projects.GET("/:id", m.Handler.GetProjectByID)
		}

		socialMedia := public.Group("/social-media")
		{
			socialMedia.GET("", m.Handler.GetSocialMedia)
			socialMedia.GET("/:id", m.Handler.GetSocialMediaByID)
		}

		tools := public.Group("/tools")
		{
			tools.GET("", m.Handler.GetTools)
			tools.GET("/:id", m.Handler.GetToolByID)
		}

		experiences := public.Group("/experiences")
		{
			experiences.GET("", m.Handler.GetExperiences)
			experiences.GET("/:id", m.Handler.GetExperienceByID)
		}
	}
}
