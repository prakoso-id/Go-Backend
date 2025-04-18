package entity

import (
	"time"

	userEntity "go-backend/internal/modules/user/domain/entity"
	"gorm.io/gorm"
)

type Tool struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"not null"`
	Icon        string         `json:"icon"`
	Category    string         `json:"category" gorm:"not null"`
	Description string         `json:"description"`
	UserID      uint           `json:"user_id" gorm:"not null"`
	User        userEntity.User `json:"user" gorm:"foreignKey:UserID"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}