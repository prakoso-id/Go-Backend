package entity

import (
	"time"

	userEntity "go-backend/internal/modules/user/domain/entity"
	"gorm.io/gorm"
)

type SocialMedia struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Platform  string         `json:"platform" gorm:"not null"`
	Url       string         `json:"url" gorm:"not null"`
	ProfileID uint           `json:"profile_id" gorm:"not null"`
	UserID    uint           `json:"user_id" gorm:"not null"`
	User      userEntity.User `json:"user" gorm:"foreignKey:UserID"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}