package entity

import (
	"time"

	"go-backend/internal/modules/user/domain/entity"
	"gorm.io/gorm"
)

type Post struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Title     string         `json:"title" gorm:"not null"`
	Content   string         `json:"content" gorm:"not null"`
	UserID    uint          `json:"user_id" gorm:"not null"`
	User      entity.User    `json:"user" gorm:"foreignKey:UserID"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}