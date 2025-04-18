package database

import (
	experienceEntity "go-backend/internal/modules/experience/domain/entity"
	imageEntity "go-backend/internal/modules/images/domain/entity"
	postEntity "go-backend/internal/modules/post/domain/entity"
	profileEntity "go-backend/internal/modules/profile/domain/entity"
	projectEntity "go-backend/internal/modules/project/domain/entity"
	socialMediaEntity "go-backend/internal/modules/socialmedia/domain/entity"
	toolEntity "go-backend/internal/modules/tool/domain/entity"
	userEntity "go-backend/internal/modules/user/domain/entity"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&userEntity.User{},
		&postEntity.Post{},
		&imageEntity.Images{},
		&projectEntity.Project{},
		&toolEntity.Tool{},
		&profileEntity.Profile{},
		&socialMediaEntity.SocialMedia{},
		&experienceEntity.Experience{},
	)
}
