package portfolio

import (
	experienceService "go-backend/internal/modules/experience/domain/service"
	postService "go-backend/internal/modules/post/domain/service"
	profileService "go-backend/internal/modules/profile/domain/service"
	projectService "go-backend/internal/modules/project/domain/service"
	"go-backend/internal/modules/portfolio/domain/service"
	"go-backend/internal/modules/portfolio/handlers"
	socialMediaService "go-backend/internal/modules/socialmedia/domain/service"
	toolService "go-backend/internal/modules/tool/domain/service"
	"gorm.io/gorm"
)

type Module struct {
	Handler *handlers.PortfolioHandler
}

func NewModule(
	db *gorm.DB,
	profileSvc profileService.ProfileService,
	postSvc postService.PostService,
	projectSvc projectService.ProjectService,
	socialMediaSvc socialMediaService.SocialMediaService,
	toolSvc toolService.ToolService,
	experienceSvc experienceService.ExperienceService,
) *Module {
	svc := service.NewPortfolioService(
		profileSvc,
		postSvc,
		projectSvc,
		socialMediaSvc,
		toolSvc,
		experienceSvc,
	)
	handler := handlers.NewPortfolioHandler(svc)

	return &Module{
		Handler: handler,
	}
}
