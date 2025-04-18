package entity

import (
	"go-backend/internal/modules/user/domain/entity"
	"time"

	"gorm.io/gorm"
)

type Experience struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Title       string         `gorm:"not null" json:"title"`
	Company     string         `gorm:"not null" json:"company"`
	Location    string         `json:"location"`
	StartDate   time.Time      `json:"start_date"`
	EndDate     *time.Time     `json:"end_date"` // Pointer to allow null for current jobs
	Description string         `json:"description"`
	TechStack   string         `gorm:"type:json" json:"tech_stack"` // Stored as JSON string
	UserID      uint           `gorm:"not null" json:"user_id"`
	User        entity.User    `gorm:"foreignKey:UserID" json:"-"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
