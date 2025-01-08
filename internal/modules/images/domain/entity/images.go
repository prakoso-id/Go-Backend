package entity

import (
	"time"

	"gorm.io/gorm"
)

type Images struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	URL       string         `json:"url" gorm:"not null"`
	PostID    uint           `json:"post_id" gorm:"not null"`
	UserID    uint           `json:"user_id" gorm:"not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
