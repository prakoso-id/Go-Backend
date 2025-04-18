package project

import (
	"go-backend/internal/modules/project/domain/repository"
	"go-backend/internal/modules/project/domain/service"
	"go-backend/internal/modules/project/handlers"
	"gorm.io/gorm"
)

type Module struct {
	Handler *handlers.ProjectHandler
}

func NewModule(db *gorm.DB) *Module {
	repo := repository.NewProjectRepository(db)
	svc := service.NewProjectService(repo)
	handler := handlers.NewProjectHandler(svc, db)

	return &Module{
		Handler: handler,
	}
}