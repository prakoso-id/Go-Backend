package profile

import (
	"go-backend/internal/modules/profile/domain/repository"
	"go-backend/internal/modules/profile/domain/service"
	"go-backend/internal/modules/profile/handlers"
	"gorm.io/gorm"
)

type Module struct {
	Handler *handlers.ProfileHandler
}

func NewModule(db *gorm.DB) *Module {
	repo := repository.NewProfileRepository(db)
	svc := service.NewProfileService(repo)
	handler := handlers.NewProfileHandler(svc)

	return &Module{
		Handler: handler,
	}
}