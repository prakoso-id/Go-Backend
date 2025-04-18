package experience

import (
	"go-backend/internal/modules/experience/domain/repository"
	"go-backend/internal/modules/experience/domain/service"
	"go-backend/internal/modules/experience/handlers"
	"gorm.io/gorm"
)

type Module struct {
	Handler *handlers.ExperienceHandler
}

func NewModule(db *gorm.DB) *Module {
	repo := repository.NewExperienceRepository(db)
	svc := service.NewExperienceService(repo)
	handler := handlers.NewExperienceHandler(svc)

	return &Module{
		Handler: handler,
	}
}
