package repository

import (
	"go-backend/internal/modules/tool/domain/entity"
	"gorm.io/gorm"
)

type ToolRepository interface {
	Create(tool *entity.Tool) error
	GetByID(id uint) (*entity.Tool, error)
	GetAll() ([]entity.Tool, error)
	Update(tool *entity.Tool) error
	Delete(id uint) error
	GetByUserID(userID uint) ([]entity.Tool, error)
}

type toolRepository struct {
	db *gorm.DB
}

func NewToolRepository(db *gorm.DB) ToolRepository {
	return &toolRepository{db: db}
}

func (r *toolRepository) Create(tool *entity.Tool) error {
	return r.db.Create(tool).Error
}

func (r *toolRepository) GetByID(id uint) (*entity.Tool, error) {
	var tool entity.Tool
	err := r.db.Preload("User").First(&tool, id).Error
	return &tool, err
}

func (r *toolRepository) GetAll() ([]entity.Tool, error) {
	var tools []entity.Tool
	err := r.db.Preload("User").Find(&tools).Error
	return tools, err
}

func (r *toolRepository) Update(tool *entity.Tool) error {
	return r.db.Save(tool).Error
}

func (r *toolRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Tool{}, id).Error
}

func (r *toolRepository) GetByUserID(userID uint) ([]entity.Tool, error) {
	var tools []entity.Tool
	err := r.db.Preload("User").Where("user_id = ?", userID).Find(&tools).Error
	return tools, err
}