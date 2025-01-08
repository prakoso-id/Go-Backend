package entity

import (
	"time"

	imageEntity "go-backend/internal/modules/images/domain/entity"
	userEntity "go-backend/internal/modules/user/domain/entity"

	"gorm.io/gorm"
)

type Post struct {
	ID        uint              `json:"id" gorm:"primaryKey"`
	Title     string            `json:"title" gorm:"not null"`
	Content   string            `json:"content" gorm:"not null"`
	UserID    uint              `json:"user_id" gorm:"not null"`
	User      userEntity.User   `json:"user" gorm:"foreignKey:UserID"`
	Images    []imageEntity.Images `json:"images" gorm:"foreignKey:PostID"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
	DeletedAt gorm.DeletedAt    `json:"-" gorm:"index"`
}
