package public

import (
	"go-backend/internal/modules/public/handlers"
	"gorm.io/gorm"
)

type Module struct {
	Handler *handlers.PublicHandler
}

func NewModule(db *gorm.DB) *Module {
	handler := handlers.NewPublicHandler(db)

	return &Module{
		Handler: handler,
	}
}
