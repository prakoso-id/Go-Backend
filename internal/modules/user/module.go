package user

import (
	"github.com/gin-gonic/gin"
	"go-backend/internal/modules/user/seeder"
	"gorm.io/gorm"
)

type Module struct {
	db *gorm.DB
}

func NewModule(db *gorm.DB) *Module {
	// Seed admin users when module is initialized
	if err := seeder.SeedAdminUsers(db); err != nil {
		panic("Failed to seed admin users: " + err.Error())
	}
	return &Module{db: db}
}

func (m *Module) RegisterRoutes(router *gin.RouterGroup) {
	RegisterRoutes(router, m.db)
}