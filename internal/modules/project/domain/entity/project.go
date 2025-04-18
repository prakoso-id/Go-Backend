package entity

import (
	"time"

	imageEntity "go-backend/internal/modules/images/domain/entity"
	userEntity "go-backend/internal/modules/user/domain/entity"

	"gorm.io/gorm"
)

type Project struct {
	ID          uint                 `json:"id" gorm:"primaryKey"`
	Name        string               `json:"name" gorm:"not null"`
	Description string               `json:"description"`
	Url         string               `json:"url"`
	UserID      uint                 `json:"user_id" gorm:"not null"`
	User        userEntity.User      `json:"user" gorm:"foreignKey:UserID"`
	Images      []imageEntity.Images `json:"images" gorm:"foreignKey:ProjectID"`
	CreatedAt   time.Time            `json:"created_at"`
	UpdatedAt   time.Time            `json:"updated_at"`
	DeletedAt   gorm.DeletedAt       `json:"-" gorm:"index"`
}
