package database

import (
	imageEntity "go-backend/internal/modules/images/domain/entity"
	postEntity "go-backend/internal/modules/post/domain/entity"
	userEntity "go-backend/internal/modules/user/domain/entity"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&userEntity.User{},
		&postEntity.Post{},
		&imageEntity.Images{},
	)
}
