package entity

import (
	"time"

	socialMediaEntity "go-backend/internal/modules/socialmedia/domain/entity"
	userEntity "go-backend/internal/modules/user/domain/entity"
	"gorm.io/gorm"
)

type Profile struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	Name         string         `json:"name" gorm:"not null"`
	Bio          string         `json:"bio"`
	ProfileImage string         `json:"profile_image"`
	Email        string         `json:"email"`
	Phone        string         `json:"phone"`
	Location     string         `json:"location"`
	UserID       uint                            `json:"user_id" gorm:"not null;uniqueIndex"`
	User         userEntity.User                  `json:"user" gorm:"foreignKey:UserID"`
	SocialMedia  []socialMediaEntity.SocialMedia `json:"social_media" gorm:"foreignKey:ProfileID"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"`
}