package socialmedia

import (
	"go-backend/internal/modules/socialmedia/domain/repository"
	"go-backend/internal/modules/socialmedia/domain/service"
	"go-backend/internal/modules/socialmedia/handlers"
	"gorm.io/gorm"
)

type Module struct {
	Handler *handlers.SocialMediaHandler
}

func NewModule(db *gorm.DB) *Module {
	repo := repository.NewSocialMediaRepository(db)
	svc := service.NewSocialMediaService(repo)
	handler := handlers.NewSocialMediaHandler(svc)

	return &Module{
		Handler: handler,
	}
}