package tool

import (
	"go-backend/internal/modules/tool/domain/repository"
	"go-backend/internal/modules/tool/domain/service"
	"go-backend/internal/modules/tool/handlers"
	"gorm.io/gorm"
)

type Module struct {
	Handler *handlers.ToolHandler
}

func NewModule(db *gorm.DB) *Module {
	repo := repository.NewToolRepository(db)
	svc := service.NewToolService(repo)
	handler := handlers.NewToolHandler(svc)

	return &Module{
		Handler: handler,
	}
}