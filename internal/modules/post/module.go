package post

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Module struct {
	db *gorm.DB
}

func NewModule(db *gorm.DB) *Module {
	return &Module{db: db}
}

func (m *Module) RegisterRoutes(router *gin.RouterGroup) {
	RegisterRoutes(router, m.db)
}