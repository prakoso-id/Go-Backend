package migrations

import (
	"go-backend/internal/modules/images/domain/entity"

	"gorm.io/gorm"
)

func CreateImagesTable(db *gorm.DB) error {
	return db.AutoMigrate(&entity.Images{})
}
